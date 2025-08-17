package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func newDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./blogs.db")
	if err != nil {
		return nil, fmt.Errorf("error opening db : %v", err)
	}
	return db, nil
}
