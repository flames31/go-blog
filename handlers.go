package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

func handleIndex() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		})
}

func handleLogin(tmpl *template.Template, db *sql.DB) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			if r.Method == http.MethodGet {
				tmpl.ExecuteTemplate(w, "login.html", LoginPageData{Title: "Login"})
				return
			}
			username := r.FormValue("username")
			password := r.FormValue("password")

			hashedPassword, err := hashPassword(password)
			if err != nil {
				log.Printf("%v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if user, err := userExists(db, username); err == nil {
				log.Printf("%v", err)
				err := checkPassword(hashedPassword, user.HashedPassword)
				if err != nil {
					log.Printf("%v", err)
					w.WriteHeader(http.StatusForbidden)
					return
				}
			} else {
				err := createUser(db, username, hashedPassword)
				if err != nil {
					log.Printf("%v", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			http.Redirect(w, r, "/blogs", http.StatusSeeOther)
		})
}

func handleGetAllBlogs(tmpl *template.Template, db *sql.DB) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			blogs, err := getAllBlogs(db)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, blogs)
		})
}
