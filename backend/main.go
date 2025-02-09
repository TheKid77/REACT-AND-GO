package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Define a simple struct for API responses
type Response struct {
	Message string `json:"message"`
}

func main() {
	r := chi.NewRouter()

	// Middleware stack
	r.Use(middleware.Logger)   // Log API requests
	r.Use(middleware.Recoverer) // Recover from panics

    	// Enable CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Adjust to match your frontend URL
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
        MaxAge:           300, // Cache preflight response for 5 minutes
	}))

	// Define routes
	r.Get("/main", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*") 
		json.NewEncoder(w).Encode(Response{Message: "Welcome to the App Bar in Andy's REACT/Golang App"})
	})

	r.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*") 
        json.NewEncoder(w).Encode(Response{Message: "Hello, " + name})
	})

	// Start the server
	http.ListenAndServe(":8080", r)
}