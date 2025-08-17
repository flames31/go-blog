package main

type Blog struct {
	Title   string
	Author  string
	Date    string
	Content string
}

type BlogPage struct {
	Title     string
	BlogTitle string
	Posts     []Blog
}

type LoginPageData struct {
	Title string
	Error string
}
