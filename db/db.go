package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DatabaseName string = "db/db.sqlite3"
var DatabaseInitializer string = "db/db-setup.sql"
var DatabaseDriver string = "sqlite"

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func GetConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect(DatabaseDriver, DatabaseName)
	checkError(err)
	return db, nil
}
