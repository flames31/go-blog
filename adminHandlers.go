package main

import (
	"io"
	"net/http"
)

func getAdminRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	io.WriteString(w, "ok")
}
