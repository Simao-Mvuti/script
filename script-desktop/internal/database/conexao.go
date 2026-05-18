package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Conectar(caminho string) *sqlx.DB {

	err := os.MkdirAll("./storage", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open("sqlite3", caminho)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err := criarTabelas(db); err != nil {
		log.Fatal(err)
	}

	return db
}

func criarTabelas(db *sqlx.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS usuarios(
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Nome VARCHAR(30) NOT NULL,
			Email TEXT NOT NULL,
			Senha TEXT NOT NULL
		)
	`

	_, err := db.Exec(query)
	return err
}
