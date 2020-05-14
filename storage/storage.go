package storage

import (
	badger "github.com/dgraph-io/badger/v2"
	log "github.com/sirupsen/logrus"
)

type Storage interface {
	UpsertMeta(fileName string, meta interface{})
	UpsertFeatures(fileName string, features []float32)
}

type ProtoStorage struct {
	db *badger.DB
}

func NewProtoStorage(dbPath string) *ProtoStorage {
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		log.WithError(err).Fatal("Cannot open data base")
	}
	defer db.Close()
	return &ProtoStorage{
		db: db,
	}
}

func (ps *ProtoStorage) Append(fileName string, features []float32) {
	err := ps.db.Update(func(txn *badger.Txn) error {
		// err := txn.Set([]byte(fileName), []byte(features))
		// return err
		return nil
	})
	log.WithError(err).WithField("fileName", fileName).Error("cannot append data")
}

func (ps *ProtoStorage) Update() {

}

func (ps *ProtoStorage) Save() error {
	return nil
}

func (ps *ProtoStorage) Load() error {
	return nil
}
