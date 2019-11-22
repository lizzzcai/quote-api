package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	DRIVER = "sqlite3"
	SQLITE_DB_LOCATION = "./quote.sqlite"
)

func ExecDB(sqlStatement string, args ...interface{}) (sql.Result, error) {
	db := dbconnect()
	defer db.Close()

	result, err := db.Exec(sqlStatement, args...)
	return result, err
}

func QueryDB(sqlStatement string, args ...interface{}) (*sql.Rows, error) {
	db := dbconnect()
	defer db.Close()

	rows, err := db.Query(sqlStatement, args...)
	return rows, err
}

func dbconnect() *sql.DB {
	db, err := sql.Open(DRIVER, SQLITE_DB_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
	return db
}