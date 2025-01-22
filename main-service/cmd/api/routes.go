package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// CORS configuration to allow domain and local development environment
	allowedOrigins := []string{
		"http://localhost:3000", // Allow requests from localhost
	}

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Cache preflight response for 5 minutes
	}))

	// Middleware
	mux.Use(middleware.Heartbeat("/ping")) // Health check endpoint
	mux.Use(middleware.Recoverer)          // Recover from panics gracefully
	mux.Use(middleware.Logger)             // Log all requests

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Resource not found", http.StatusNotFound)
	})

	mux.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	mux.Use(httprate.LimitByIP(100, 1*time.Minute)) // 100 requests per IP per minute

	// API Routes
	mux.Route("/api/v1", func(r chi.Router) {
		r.Post("/main-service/save-image-data", app.saveImageDataHandler)
	})
	return mux
}
