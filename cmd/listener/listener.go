package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Define the HTTP request handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Set the response content type
		w.Header().Set("Content-Type", "text/plain")

		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		// Print the request body
		fmt.Println("Request Body:", string(body))
	}

	// Register the handler function to the default HTTP server
	http.HandleFunc("/", handler)

	// Start the HTTP server and listen on port 8080
	log.Println("Server started on http://localhost:5555")
	log.Fatal(http.ListenAndServe(":5555", nil))
}
