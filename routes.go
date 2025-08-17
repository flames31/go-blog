package main

import (
	"html/template"

	"github.com/gorilla/mux"
)

func addRoutes(mux *mux.Router, tmpl *template.Template) {

	mux.Handle("/login", handleLogin(tmpl))
	mux.Handle("/blogs", handleGetAllBlogs(tmpl))
	mux.Handle("/", handleIndex())
}
