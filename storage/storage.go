package storage

import (
	"fmt"
	badger "github.com/dgraph-io/badger/v2"
	log "github.com/sirupsen/logrus"
)

type Storage struct {
	db *badger.DB
}

func NewStorage(dbPath string) *Storage {
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		log.WithError(err).Fatal("Cannot open data base")
	}
	return &Storage{
		db: db,
	}
}

func NewInMemStorage() *Storage {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		log.WithError(err).Fatal("Cannot open data base")
	}
	return &Storage{
		db: db,
	}
}

func (s* Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) GetFiles() ([]string, error) {
	var files []string

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		iter := txn.NewIterator(opts)
		defer iter.Close()
		for iter.Rewind(); iter.Valid(); iter.Next() {
			key := iter.Item().Key()
			files = append(files, string(key))
		}
		return nil
	})

	return files, err
}

func (s *Storage) UpsertFeatures(fileName string, features []float32) error {
	err := s.db.Update(func(txn *badger.Txn) error {
		data, err := MarshalFloats(features)
		if err != nil {
			return fmt.Errorf("cannot save features for file %s: %v", fileName, err)
		}
		err = txn.Set([]byte(fileName), data)
		return err
	})
	return err
}

func (s* Storage) GetFeatures(fileName string) ([]float32, error) {
	var features []float32
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(fileName))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		features, err = UnmarshallFloats(val)
		if err != nil {
			return err
		}
		return nil
	})

	return features, err
}
