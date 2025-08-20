package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func createUser(db *sql.DB, username, hashedPassword string) error {
	_, err := db.Exec(`INSERT INTO USERS VALUES (?, ?, ?, ?, ?)`, uuid.New(), time.Now(), time.Now(), username, hashedPassword)
	return err
}

func getUser(db *sql.DB, username string) (User, error) {
	var u User
	row := db.QueryRow(`SELECT * FROM users WHERE username = ?;`, username)

	err := row.Scan(
		&u.ID,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Username,
		&u.HashedPassword,
	)

	if err != nil {
		return User{}, err
	}

	return u, nil
}
