package main

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func startServer() error {

	db, err := newDB()
	if err != nil {
		return fmt.Errorf("failed to open db : %v", err)
	}
	defer db.Close()

	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	srv := NewServer(tmpl, db)
	httpSrv := &http.Server{
		Addr:         ":42069",
		Handler:      srv,
		IdleTimeout:  time.Minute * 1,
		ReadTimeout:  time.Minute * 1,
		WriteTimeout: time.Minute * 2,
	}

	log.Printf("Server is listening on port : %v", httpSrv.Addr)
	err = httpSrv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server is shutting down")
	} else {
		return err
	}

	return nil
}

func NewServer(tmpl *template.Template, db *sql.DB) http.Handler {
	mux := mux.NewRouter()
	addRoutes(mux, tmpl, db)
	var handler http.Handler = mux

	handler = logRequest(handler)

	return handler
}
