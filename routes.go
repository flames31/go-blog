package main

import (
	"database/sql"
	"html/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func addRoutes(mux *mux.Router, tmpl *template.Template, db *sql.DB, store *sessions.CookieStore) {

	mux.Handle("/login", handleLogin(tmpl, db, store))
	mux.Handle("/blogs", isAuthenticated(handleGetBlogs(tmpl, db), store)).Methods("GET")
	mux.Handle("/blogs", isAuthenticated(handlePostBlog(tmpl, db, store), store)).Methods("POST")
	mux.Handle("/", handleIndex())
	mux.Handle("/blogs/new", isAuthenticated(handleCreateBlog(tmpl, db), store))
	mux.Handle("/myblogs", isAuthenticated(handleMyBlogs(tmpl, db, store), store)).Methods("GET")
	mux.Handle("/blogs/edit/{id}", isAuthenticated(handleEditBlog(tmpl, db, store), store))
	mux.Handle("/blogs/delete/{id}", isAuthenticated(handleDeleteBlog(tmpl, db, store), store))
}
