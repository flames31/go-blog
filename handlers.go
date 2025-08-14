package main

import "net/http"

func getMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/posts", allPosts)
	return logRequest(mux)
}
