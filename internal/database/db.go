package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/syeadz/classup/internal/settings"
)

func GetDb() *sql.DB {
	db, err := sql.Open("sqlite3", settings.GetConfig().DbPath)

	if err != nil {
		log.Fatalln("Couldn't open")
	}

	return db
}
