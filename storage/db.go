package storage

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitDb(filePath string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite", filePath)
	if err != nil {
		return
	}

	// enable foreign keys
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return
	}

	return
}
