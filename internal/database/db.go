package database

import (
	"log"
	"os"

	bolt "go.etcd.io/bbolt"
)

func Test() *bolt.DB {
	db, err := bolt.Open(os.Getenv("HOME")+"/.classup/test.db", 0666, nil)

	if err != nil {
		log.Println(err)
	}

	return db
}
