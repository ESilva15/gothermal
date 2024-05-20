package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DBAccess interface {
	// I should make a query struct so that I can make generic structs
	// And then slap them into different types of DBs
	createq(query string) error
	selectq(query string) error
	deleteq(query string) error
	updateq(query string) error
}

type Database struct {
	DB   *sql.DB
	Path string
}

func (s *Database) Initialize(path string) error {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	s.Path = path
	s.DB = db

	return nil
}
