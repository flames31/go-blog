package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func getAllBlogs(db *sql.DB) ([]Blog, error) {
	var blogs []Blog

	rows, err := db.Query(`SELECT * FROM blogs;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b Blog
		if err = rows.Scan(
			&b.ID,
			&b.Title,
			&b.Author,
			&b.Date,
			&b.Content,
			&b.UserID,
		); err != nil {
			return nil, err
		}

		blogs = append(blogs, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func postBlog(db *sql.DB, title, author, content, userID string) error {
	_, err := db.Exec(`INSERT INTO blogs values (?, ?, ?, ?, ?, ?)`, uuid.New(), title, author, time.Now(), content, userID)
	return err
}

func getBlogByUserID(db *sql.DB, userID uuid.UUID) ([]Blog, error) {
	var blogs []Blog

	rows, err := db.Query(`SELECT * FROM blogs WHERE userID = ?;`, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b Blog
		if err = rows.Scan(
			&b.ID,
			&b.Title,
			&b.Author,
			&b.Date,
			&b.Content,
			&b.UserID,
		); err != nil {
			return nil, err
		}

		blogs = append(blogs, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func updateBlog(db *sql.DB, blogID uuid.UUID, title, author, content string) error {
	_, err := db.Exec(`UPDATE blogs SET title = ?, author = ?, content = ? WHERE ID = ?;`, title, author, content, blogID)
	return err
}

func getBlogByID(db *sql.DB, blogID uuid.UUID) (Blog, error) {
	var blog Blog

	row := db.QueryRow(`SELECT * FROM blogs WHERE id = ?;`, blogID)
	err := row.Scan(
		&blog.ID,
		&blog.Title,
		&blog.Author,
		&blog.Date,
		&blog.Content,
		&blog.UserID,
	)
	if err != nil {
		return Blog{}, err
	}
	return blog, nil
}

func deleteBlog(db *sql.DB, blogID uuid.UUID) error {
	_, err := db.Exec(`DELETE FROM blogs WHERE ID = ?;`, blogID)
	return err
}
