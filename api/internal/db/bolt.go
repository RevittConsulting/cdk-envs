package db

import (
	"github.com/boltdb/bolt"
	"log"
)

type Db struct {
	Db *bolt.DB
}

func New(filePath string) *Db {
	return &Db{
		Db: openDB(filePath),
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
	return b.Db.Close()
}
