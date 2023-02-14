package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	settings "github.com/syeadz/classup/configs"
)

func GetDb() *sql.DB {
	db, err := sql.Open("sqlite3", settings.GetConfig().DbPath)

	if err != nil {
		log.Fatalln("Couldn't open")
	}

	return db
}
