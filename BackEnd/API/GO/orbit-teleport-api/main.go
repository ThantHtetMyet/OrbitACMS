package main

import (
	"log"
	"net/http"

	"orbit-teleport-api/internal/router"
)

func main() {
	mux := router.New()

	addr := ":8080"
	log.Printf("API server running on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
