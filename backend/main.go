package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Data structure to serve as a response
type Response struct {
    Data []string `json:"data"`
}

func main() {
    r := chi.NewRouter()

    // Middleware
    r.Use(middleware.Logger)  // Logs each HTTP request
    r.Use(middleware.Recoverer) // Graceful panic recovery

    // Define routes
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        response := Response{Data: []string{"Hello, World!"}}
        json.NewEncoder(w).Encode(response)
    })

    r.Get("/data", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        response := Response{Data: []string{"item1", "item2", "item3"}}
        json.NewEncoder(w).Encode(response)
    })

    // Start the server
    http.ListenAndServe(":8080", r)
}