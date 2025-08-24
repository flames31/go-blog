package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

func logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		h.ServeHTTP(w, r)
		t2 := time.Since(t1)
		log.Printf("%-10v %-10v %-10v", r.Method, r.URL, t2)
	})
}

func isAuthenticated(h http.Handler, store *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "userSessions")
		userIDStr, ok := session.Values["user_id"].(string)
		if !ok || userIDStr == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		h.ServeHTTP(w, r)
	})
}
