package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Import the pq package for PostgreSQL
)

// Config holds configuration values, including database connections.
type Config struct {
	DB *sql.DB // Database connection
}

type ImageData struct {
	ImageURL     string `json:"imageUrl"`
	CreationTime string `json:"creationTime"`
}

func (app *Config) saveImageDataHandler(w http.ResponseWriter, r *http.Request) {
	var data ImageData

	// Decode the request body into the ImageData struct
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Here you would save the image data to your database or process it
	log.Printf("Received image data: %v", data)

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Image data saved successfully!"})
}
