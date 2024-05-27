package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define a handler function for the /hello route
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        // Set the response content type
        w.Header().Set("Content-Type", "text/plain")

        // Write the response body
        fmt.Fprintf(w, "Hello, World!")
    })

    // Start the HTTP server on port 8080
    fmt.Println("Server listening on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
