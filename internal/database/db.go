package database

import (
	"log"
	"os"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/codec/protobuf"
)

func Test() {
	db, err := storm.Open(os.Getenv("HOME")+"/.classup/test.db", storm.Codec(protobuf.Codec))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
