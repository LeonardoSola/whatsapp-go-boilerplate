package sqlite

import (
	"database/sql"

	"github.com/LeonardoSola/whatsapp-go-boilerplate/adapters/env"

	_ "github.com/mattn/go-sqlite3"
)

var SQLiteDB *sql.DB

func init() {
	dbInstance, err := sql.Open("sqlite3", env.SQLiteDB)
	if err != nil {
		panic(err)
	}

	SQLiteDB = dbInstance
}
