package main

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	ID      uuid.UUID
	Title   string
	Author  string
	Date    time.Time
	Content string
	UserID  uuid.UUID
}

type BlogPage struct {
	Title     string
	BlogTitle string
	Blogs     []Blog
}

type LoginPageData struct {
	Title string
	Error string
}

type User struct {
	ID             uuid.UUID
	Username       string
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
