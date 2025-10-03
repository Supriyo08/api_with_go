package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Register a handler for the "/hello" route.
	http.HandleFunc("/hello", helloHandler)

	// Start the server on port 8080.
	fmt.Println("Server is listening on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

// helloHandler is the handler function for the "/hello" route.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET.
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Write a response to the client.
	fmt.Fprintf(w, "Hello, Go API!")
}