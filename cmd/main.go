package main

import (
	"log"
)

func main() {
	srv, err := InitializeServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	srv.Start()
}
