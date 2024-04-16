package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// GetDBConnection returns a connection to the SQLite database.
func GetDBConnection() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Enable logging (optional)
	db.LogMode(true)

	// Migrate the database schema
	db.AutoMigrate(&Admission{})

	return db, nil
}
