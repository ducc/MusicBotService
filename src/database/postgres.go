package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const BASE_URL = "%s://%s:%s@%s/%s"

type Database struct {
	driver string
	url    string
	db     *sql.DB
}

func NewDatabase(driver, url string) *Database {
	db := new(Database)
	db.driver = driver
	db.url = url
	return db
}

func (db *Database) Open() (*Database, error) {
	database, err := sql.Open(db.driver, db.url)
	if err != nil {
		return nil, err
	}
	db.db = database
	return db, nil
}

func (db *Database) Close() error {
	err := db.db.Close()
	return err
}

func (db Database) Prepare(sql string) *sql.Stmt {
	statement, err := db.db.Prepare(sql)
	if err != nil {
		fmt.Println("Error preparing statement,", err)
		return nil
	}
	return statement
}

func (db Database) GetDB() *sql.DB {
	return db.db
}
