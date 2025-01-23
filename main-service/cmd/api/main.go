// Main Service: Handles requests from the React Web App
package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	// Service configurations
	ServicePort = "8080"         // Port on which the service will run
	ServiceName = "Main-Service" // Name of the service

)

func main() {
	// Log startup message
	log.Printf("Starting %s on port %s...", ServiceName, ServicePort)

	// Initialize the GORM DB connection
	db, err := NewDBConnection()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Create a Config instance
	app := &Config{DB: db}

	// Create and configure the HTTP server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", ServicePort),
		Handler: app.routes(),
	}

	// Start the server
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
