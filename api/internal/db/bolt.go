package db

import (
	"github.com/boltdb/bolt"
	"log"
)

type Db struct {
	db *bolt.DB
}

func New(filePath string) *Db {
	return &Db{
		db: openDB(filePath),
	}
}

func openDB(filePath string) *bolt.DB {
	db, err := bolt.Open(filePath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (b *Db) Close() error {
	return b.db.Close()
}
