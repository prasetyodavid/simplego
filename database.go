package database

import (
	"gorm.io/gorm"
)

func connection() {
	db, err = gorm.Open("sqlite3", "database.db")

	if err != nil {
		panic("Failed to open the SQLite database.")
	}

	return db, err
}
