package main

import (
	route "github.com/domyid/domyapi/route"
	"log"
	"net/http"
)

func main() {
	// Handle the WebHook route
	http.HandleFunc("/webhook", route.URL)

	// Start the server on a specific port, e.g., 8080
	port := ":80"
	log.Printf("Serber running on port: ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
