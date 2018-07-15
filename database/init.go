package database

import (
	"database/sql"
	"log"
)

var globalDB *sql.DB

func Init(driver, connString string) error {
	db, err := sql.Open(driver, connString)
	if err != nil {
		return err
	}

	globalDB = db

	if err := globalDB.Ping(); err != nil {
		return err
	}

	log.Println("Database connected successfully")
	return nil
}

func Get() *sql.DB {
	return globalDB
}
