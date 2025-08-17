package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	tmpl *template.Template
}

func startServer() error {
	srv := NewServer()
	httpSrv := &http.Server{
		Addr:         ":42069",
		Handler:      srv,
		IdleTimeout:  time.Minute * 1,
		ReadTimeout:  time.Minute * 1,
		WriteTimeout: time.Minute * 2,
	}

	log.Printf("Server is listening on port : %v", httpSrv.Addr)
	err := httpSrv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server is shutting down")
	} else {
		return err
	}

	return nil
}

func NewServer() http.Handler {
	srv := &server{
		tmpl: template.Must(template.ParseGlob("templates/*.html")),
	}

	mux := mux.NewRouter()
	addRoutes(mux, srv.tmpl)
	var handler http.Handler = mux

	handler = logRequest(handler)

	return handler
}
