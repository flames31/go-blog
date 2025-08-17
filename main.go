package main

import (
	"log"
	"os"
)

func main() {
	err := startServer()
	if err != nil {
		log.Printf("error with server : %v", err)
		os.Exit(1)
	}
}
