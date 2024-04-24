package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define our router, we'll use standard library
	r := http.NewServeMux()

	// Register our routes (endpoint handlers)
	r.HandleFunc("/", helloWorld)

	// Start our webserver
	fmt.Println("Server running!")
	http.ListenAndServe(":8080", r)
}

// Hello World endpoint handler
func helloWorld(w http.ResponseWriter, r *http.Request) {
	// Respond with some text
	w.Write([]byte("Hello World!"))
}
