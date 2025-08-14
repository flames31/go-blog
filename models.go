package main

type Post struct {
	Title   string
	Author  string
	Date    string
	Content string
}

type BlogPage struct {
	Title     string
	BlogTitle string
	Posts     []Post
}

type LoginPageData struct {
	Title string
	Error string
}
