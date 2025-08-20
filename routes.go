package main

import (
	"database/sql"
	"html/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func addRoutes(mux *mux.Router, tmpl *template.Template, db *sql.DB, store *sessions.CookieStore) {

	mux.Handle("/login", handleLogin(tmpl, db, store))
	mux.Handle("/blogs", handleGetBlogs(tmpl, db)).Methods("GET")
	mux.Handle("/blogs", handlePostBlog(tmpl, db, store)).Methods("POST")
	mux.Handle("/", handleIndex())
	mux.Handle("/blogs/new", handleCreateBlog(tmpl, db))
	mux.Handle("/myblogs", handleMyBlogs(tmpl, db, store)).Methods("GET")
	mux.Handle("/blogs/edit/{id}", handleEditBlog(tmpl, db, store))
	mux.Handle("/blogs/delete/{id}", handleDeleteBlog(tmpl, db, store))
}
