package main

import (
	"log"
	"net/http"
	"time"
)

func logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		h.ServeHTTP(w, r)
		t2 := time.Since(t1)
		log.Printf("%-10v %-10v %-10v", r.Method, r.URL, t2)
	})
}
