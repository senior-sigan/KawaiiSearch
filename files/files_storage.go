package files

import (
	"context"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"path"
	"path/filepath"
)

type FilesStorage struct {
	rootDir      string
	extensions   map[string]bool
	cancel       func()
	watcher      *fsnotify.Watcher
	OnNewFile    chan string
	OnDeleteFile chan string
	OnRenameFile chan string
	OnChangeFile chan string
	Errors       chan error
}

func NewFilesStorage(rootDir string) *FilesStorage {
	return &FilesStorage{
		rootDir: rootDir,
		extensions: map[string]bool{
			".jpg": true, ".png": true, ".jpeg": true, ".JPG": true, ".PNG": true, ".JPEG": true,
		},
		OnNewFile:    make(chan string),
		OnDeleteFile: make(chan string),
		OnChangeFile: make(chan string),
		OnRenameFile: make(chan string),
		Errors:       make(chan error),
	}
}

func (fs *FilesStorage) Connect(ctx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithCancel(ctx)
	fs.cancel = cancel

	go fs.start(ctx)
}

func (fs *FilesStorage) Close() error {
	if fs.cancel != nil {
		fs.cancel()
	}
	return fs.watcher.Close()
}

func (fs *FilesStorage) isAllowedExt(filename string) bool {
	ext := filepath.Ext(filename)
	if _, ok := fs.extensions[ext]; ok {
		return true
	}
	return false
}

func (fs *FilesStorage) start(ctx context.Context) {
	fs.listImages()
	fs.watch(ctx)
}

func (fs *FilesStorage) listImages() {
	logrus.WithField("rootDir", fs.rootDir).Info("Searching files")
	for ext, _ := range fs.extensions {
		pattern := path.Join(fs.rootDir, "*"+ext)
		files, err := filepath.Glob(pattern)
		if err != nil {
			fs.Errors <- err
			continue
		}
		for _, f := range files {
			logrus.WithField("file", f).WithField("event", "FILE").Info("FS listdir")
			fs.OnNewFile <- f
		}
	}
}

func (fs *FilesStorage) watch(ctx context.Context) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fs.Errors <- err
		return
	}
	fs.watcher = watcher

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if !fs.isAllowedExt(event.Name) {
					continue
				}
				logrus.WithField("file", event.Name).WithField("event", event.Op.String()).Info("FS Event")

				if event.Op&fsnotify.Create == fsnotify.Create {
					fs.OnNewFile <- event.Name
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					fs.OnDeleteFile <- event.Name
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					fs.OnRenameFile <- event.Name
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fs.OnChangeFile <- event.Name
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fs.Errors <- err
			}
		}
	}()

	err = watcher.Add(fs.rootDir)
	if err != nil {
		fs.Errors <- err
	}
}
