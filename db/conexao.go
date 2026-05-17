package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func Connect() {
	db, err := sqlx.Connect("sqlite3", "./empresa.db")
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
