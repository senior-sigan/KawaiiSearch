package worker

import (
	"github.com/sirupsen/logrus"
	"kawaii_search/files"
	"kawaii_search/storage"
)

type HashWorker struct {
	hashesStorage *storage.Storage
	fileStorage   *files.FilesStorage
}

func NewHashWorker(hashesStorage *storage.Storage, fileStorage *files.FilesStorage) *HashWorker {
	return &HashWorker{
		hashesStorage: hashesStorage,
		fileStorage:   fileStorage,
	}
}

func (w *HashWorker) calculateHash(filename string) ([]float32, error) {
	// TODO: call embedder to get hash
	return []float32{42}, nil
}

func (w *HashWorker) onNewFile(filename string) {
	contains, err := w.hashesStorage.Contains(filename)
	if err != nil {
		logrus.WithError(err).WithField("file", filename).Error("Cannot handle new file")
		return
	}
	if contains {
		// nothing to do
		return
	}

	hash, err := w.calculateHash(filename)
	if err != nil {
		logrus.WithError(err).WithField("file", filename).Error("Cannot calculate hash for the new file")
		return
	}

	err = w.hashesStorage.UpsertFeatures(filename, hash)
	if err != nil {
		logrus.WithError(err).WithField("file", filename).Error("Cannot save hash for the file")
	}
	logrus.WithField("file", filename).Info("File hash is saved")

	// TODO: update kNN index
}

func (w *HashWorker) onRenameFile(filename string) {
	logrus.WithField("file", filename).Warn("onRenameFile is not implemented")

}

func (w *HashWorker) onChangeFile(filename string) {
	hash, err := w.calculateHash(filename)
	if err != nil {
		logrus.WithError(err).WithField("file", filename).Error("Cannot calculate hash for the changed file")
		return
	}
	err = w.hashesStorage.UpsertFeatures(filename, hash)
	if err != nil {
		logrus.WithError(err).WithField("file", filename).Error("Cannot update hash for the file")
	}
	logrus.WithField("file", filename).Info("File hash is updated")
	// TODO: update kNN index
}

func (w *HashWorker) onDeleteFile(filename string) {
	err := w.hashesStorage.Remove(filename)
	if err != nil {
		logrus.WithError(err).WithField("file", filename).Error("Cannot handle deleted file")
	}
	logrus.WithField("file", filename).Info("File hash is deleted")
	// TODO: update kNN index
}

func (w *HashWorker) Start() {
	w.fileStorage.Connect(nil)

	for {
		select {
		case file, ok := <-w.fileStorage.OnNewFile:
			if !ok {
				continue
			}
			w.onNewFile(file)
		case file, ok := <-w.fileStorage.OnRenameFile:
			if !ok {
				continue
			}
			w.onRenameFile(file)
		case file, ok := <-w.fileStorage.OnChangeFile:
			if !ok {
				continue
			}
			w.onChangeFile(file)
		case file, ok := <-w.fileStorage.OnDeleteFile:
			if !ok {
				continue
			}
			w.onDeleteFile(file)
		}
	}
}
