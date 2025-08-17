package main

import "database/sql"

func getAllBlogs(db *sql.DB) ([]Blog, error) {
	blogs := make([]Blog, 0)

	rows, err := db.Query(`SELECT * FROM blogs;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b Blog
		if err = rows.Scan(&b); err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}
