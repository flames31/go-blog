package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var tmpl *template.Template

func startServer() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	mux := getMux()

	srv := &http.Server{
		Addr:         ":42069",
		Handler:      mux,
		IdleTimeout:  time.Minute * 1,
		ReadTimeout:  time.Minute * 1,
		WriteTimeout: time.Minute * 2,
	}

	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server is shutting down")
	} else {
		log.Printf("server error : %v", err)
		os.Exit(1)
	}
}
