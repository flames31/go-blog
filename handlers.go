package main

import (
	"html/template"
	"net/http"
)

func handleIndex() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		})
}

func handleLogin(tmpl *template.Template) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			if r.Method == http.MethodGet {
				tmpl.ExecuteTemplate(w, "login.html", LoginPageData{Title: "Login"})
				return
			}
			//username := r.FormValue("username")
			//password := r.FormValue("password")

			http.Redirect(w, r, "/blogs", http.StatusSeeOther)

		})
}

func handleGetAllBlogs(tmpl *template.Template) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			data := BlogPage{
				Title:     "My Blog",
				BlogTitle: "Go Blogger",
				Posts: []Post{
					{Title: "First Post", Author: "Rahul", Date: "2025-08-13", Content: "This is my first blog post in Go!"},
					{Title: "Second Post", Author: "Rahul", Date: "2025-08-14", Content: "Learning html/template is fun."},
				},
			}
			tmpl.Execute(w, data)
		})
}
