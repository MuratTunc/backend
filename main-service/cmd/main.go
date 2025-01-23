package main

import (
	"log"
	"net/http"

	kitHttp "github.com/go-kit/kit/transport/http" // Alias Go-Kit HTTP transport
	"github.com/gorilla/mux"                       // Gorilla Mux router
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"main-service/cmd/models"
	"main-service/cmd/service"
	"main-service/cmd/transport"
)

func main() {
	// Database connection setup
	dsn := "host=localhost user=yourUser password=yourPassword dbname=yourDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate schema (make sure your table is set up)
	err = db.AutoMigrate(&models.ImageData{})
	if err != nil {
		log.Fatalf("Failed to migrate schema: %v", err)
	}

	// Initialize service (passing DB connection to the service layer)
	svc := service.NewService(db)

	// Create the endpoint for saving image data
	saveImageDataEndpoint := transport.MakeSaveImageDataEndpoint(svc)

	// Create HTTP handler using Go-Kit
	saveImageDataHandler := kitHttp.NewServer(
		saveImageDataEndpoint,                // The endpoint we created
		transport.DecodeSaveImageDataRequest, // Request decoder
		transport.EncodeResponse,             // Response encoder
	)

	// Set up Gorilla Mux router
	r := mux.NewRouter()

	// Define the route for saving image data
	r.Handle("/api/v1/main-service/save-image-data", saveImageDataHandler).Methods("POST")

	// Start the HTTP server
	log.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
