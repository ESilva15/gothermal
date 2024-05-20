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

type SQLite struct {
	DB   *sql.DB
	Path string
}

func (s *SQLite) Initialize(path string) error {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	s.Path = path
	s.DB = db

	return nil
}

func (s *SQLite) Createq(query string) error {
	return nil
}

func (s *SQLite) Selectq(query string) error {
	return nil
}

func (s *SQLite) Deleteq(query string) error {
	return nil
}

func (s *SQLite) Updateq(query string) error {
	return nil
}
