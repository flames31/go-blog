package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl.ExecuteTemplate(w, "login.html", LoginPageData{Title: "Login"})
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(username + " " + password)

	tmpl.Execute(w, LoginPageData{
		Title: "Login",
	})
}

func allPosts(w http.ResponseWriter, r *http.Request) {

	data := BlogPage{Title: "My Blog",
		BlogTitle: "Go Blogger",
		Posts: []Post{
			{Title: "First Post", Author: "Rahul", Date: "2025-08-13", Content: "This is my first blog post in Go!"},
			{Title: "Second Post", Author: "Rahul", Date: "2025-08-14", Content: "Learning html/template is fun."},
		},
	}
	tmpl.Execute(w, data)
}
