package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

const keySrvAddr = "ServerAddr"

func main() {
	userMux := http.NewServeMux()
	adminMux := http.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	srv1 := &http.Server{
		Addr:    ":42069",
		Handler: userMux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keySrvAddr, l.Addr().String())
			return ctx
		},
	}

	srv2 := &http.Server{
		Addr:    ":42070",
		Handler: adminMux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keySrvAddr, l.Addr().String())
			return ctx
		},
	}
	userMux.HandleFunc("/", getUserRoot)
	adminMux.HandleFunc("/", getAdminRoot)

	go func() {
		defer cancel()
		err := srv1.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Printf("server is shutting down")
		} else {
			log.Printf("server error : %v", err)
			os.Exit(1)
		}
	}()

	go func() {
		defer cancel()
		err := srv2.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Printf("server is shutting down")
		} else {
			log.Printf("server error : %v", err)
			os.Exit(1)
		}
	}()
	<-ctx.Done()
}

func getUserRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from user : %v", r.URL)
	w.WriteHeader(200)
	io.WriteString(w, "ok")
}

func getAdminRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from admin on path : %v", r.URL)
	w.WriteHeader(200)
	io.WriteString(w, "ok")
}
