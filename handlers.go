package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func handleIndex() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		})
}

func handleLogin(tmpl *template.Template, db *sql.DB, store *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			if r.Method == http.MethodGet {
				tmpl.ExecuteTemplate(w, "login.html", LoginPageData{Title: "Login"})
				return
			}
			username := r.FormValue("username")
			password := r.FormValue("password")

			session, _ := store.Get(r, "userSessions")
			hashedPassword, err := hashPassword(password)
			if err != nil {
				log.Printf("%v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			usr, err := getUser(db, username)
			log.Printf("%v %v", usr, err)
			if user, err := getUser(db, username); err == nil {
				log.Printf("%v", user)
				err := checkPassword(user.HashedPassword, password)
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
			user, _ := getUser(db, username)
			session.Values["user_id"] = user.ID.String()
			if err := session.Save(r, w); err != nil {
				fmt.Println(err)
			}
			http.Redirect(w, r, "/blogs", http.StatusSeeOther)
		})
}

func handleGetBlogs(tmpl *template.Template, db *sql.DB) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			blogs, err := getAllBlogs(db)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = tmpl.ExecuteTemplate(w, "blogs.html", blogs)
			if err != nil {
				fmt.Println(err)
			}
		})
}
