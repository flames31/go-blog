package main

import (
	"database/sql"
	"html/template"

	"github.com/gorilla/mux"
)

func addRoutes(mux *mux.Router, tmpl *template.Template, db *sql.DB) {

	mux.Handle("/login", handleLogin(tmpl, db))
	mux.Handle("/blogs", handleGetAllBlogs(tmpl, db))
	mux.Handle("/", handleIndex())
}
