package web

import (
    "encoding/json"
    "net/http"
)

// Start starts the HTTP server
func Start() error {
    http.HandleFunc("/hello", handleHello)
    return http.ListenAndServe(":8080", nil)
}

// handleHello is the handler for the "/hello" endpoint
func handleHello(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{
        "message": "Hello, world!",
    }
    json.NewEncoder(w).Encode(data)
}
