package db

import (
	"database/sql"
	// "fmt"
	"goride/internal/store"
	"os"

	_ "github.com/shaxbee/go-spatialite"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)
func open(dbName string) (*gorm.DB, error) {

	// Create the temp directory if it doesn't exist
	err := os.MkdirAll("/tmp", 0755)
	if err != nil {
		return nil, err
	}

	// Open the SQLite database with mattn/go-sqlite3
	sqlDB, err := sql.Open("spatialite", dbName)
	if err != nil {
		return nil, err
	}


	// Wrap the raw SQL connection with GORM
	db, err := gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MustOpen(dbName string) *gorm.DB {

	if dbName == "" {
		dbName = "goth.db"
	}

	db, err := open(dbName)
	if err != nil {
		panic(err)
	}

	// Automatically migrate the database schemas for User and Session
	err = db.AutoMigrate(&store.User{}, &store.Session{})
	if err != nil {
		panic(err)
	}

	return db
}
