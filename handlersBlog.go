package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func handleCreateBlog(tmpl *template.Template, db *sql.DB) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			if r.Method == http.MethodGet {
				tmpl.ExecuteTemplate(w, "createBlog.html", nil)
				return
			}
		})
}

func handlePostBlog(tmpl *template.Template, db *sql.DB, store *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			title := r.FormValue("title")
			content := r.FormValue("content")
			author := r.FormValue("author")

			session, _ := store.Get(r, "userSessions")
			fmt.Println(session.Values)
			userID, ok := session.Values["user_id"].(string)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			err := postBlog(db, title, author, content, userID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/blogs", http.StatusSeeOther)
		})
}

func handleMyBlogs(tmpl *template.Template, db *sql.DB, store *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, "userSessions")
			userIDStr := session.Values["user_id"].(string)
			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				log.Printf("ERROR : %v", err)
				return
			}
			blogs, err := getBlogByUserID(db, userID)
			if err != nil {
				log.Printf("ERROR : %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fmt.Println(blogs)
			err = tmpl.ExecuteTemplate(w, "myblogs.html", blogs)
			if err != nil {
				fmt.Println(err)
			}
		})
}

func handleEditBlog(tmpl *template.Template, db *sql.DB, store *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		blogIDStr := mux.Vars(r)["id"]
		blogID, err := uuid.Parse(blogIDStr)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			log.Printf("ERROR : %v", err)
			return
		}
		session, _ := store.Get(r, "userSessions")
		userIDStr := session.Values["user_id"].(string)
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			log.Printf("ERROR : %v", err)
			return
		}

		if r.Method == http.MethodGet {
			blog, err := getBlogByID(db, blogID)
			if err != nil {
				http.Error(w, "Blog not found", http.StatusNotFound)
				log.Printf("ERROR : %v", err)
				return
			}

			if userID != blog.UserID {
				http.Error(w, "Only blog owner can edit posts!", http.StatusForbidden)
				return
			}
			tmpl.ExecuteTemplate(w, "blog_edit.html", blog)
			return
		}

		if r.Method == http.MethodPost {
			title := r.FormValue("title")
			author := r.FormValue("author")
			content := r.FormValue("content")

			err := updateBlog(db, blogID, title, author, content)
			if err != nil {
				http.Error(w, "Failed to update blog", http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/myblogs", http.StatusSeeOther)
		}
	})
}

func handleDeleteBlog(tmpl *template.Template, db *sql.DB, store *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			blogIDStr := mux.Vars(r)["id"]
			blogID, err := uuid.Parse(blogIDStr)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				log.Printf("ERROR : %v", err)
				return
			}
			session, _ := store.Get(r, "userSessions")
			userIDStr := session.Values["user_id"].(string)
			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				log.Printf("ERROR : %v", err)
				return
			}
			blog, err := getBlogByID(db, blogID)
			if err != nil {
				http.Error(w, "Blog not found", http.StatusNotFound)
				log.Printf("ERROR : %v", err)
				return
			}

			if userID != blog.UserID {
				http.Error(w, "Only blog owner can edit posts!", http.StatusForbidden)
				return
			}

			err = deleteBlog(db, blogID)
			if err != nil {
				log.Printf("ERROR : %v", err)
				http.Error(w, "error deleting blog", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/myblogs", http.StatusSeeOther)
		})
}
