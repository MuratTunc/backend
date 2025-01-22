package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config holds configuration values, including database connections.
type Config struct {
	DB *gorm.DB // Database connection
}

// ImageData struct represents the database table.
type ImageData struct {
	ID           uint   `gorm:"primaryKey"` // Auto-incrementing ID field
	ImageURL     string `json:"imageUrl"`
	CreationTime string `json:"creationTime"`
	Title        string `json:"title"`
	Text         string `json:"text"`
}

// NewDBConnection initializes a new GORM database connection.
func NewDBConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=yourUser password=yourPassword dbname=yourDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Migrate the schema (create the table if it doesn't exist)
	err = db.AutoMigrate(&ImageData{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate schema: %w", err)
	}

	return db, nil
}

// saveImageDataHandler processes the incoming request and saves image data with title and text
func (app *Config) saveImageDataHandler(w http.ResponseWriter, r *http.Request) {
	var data ImageData

	// Decode the request body into the ImageData struct
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log the received data (for debugging)
	log.Printf("Received image data: %v", data)

	// Save the data into the database using GORM
	result := app.DB.Create(&data)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Image data saved successfully!"})
}
