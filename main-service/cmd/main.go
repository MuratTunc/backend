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

	"github.com/fatih/color" // Import the color package
)

// Define constants for the database connection and service port
const (
	DBHost             = "localhost"
	DBUser             = "postgres"
	DBPassword         = "4258"
	DBName             = "postgres"
	DBPort             = "5439"
	SSlMode            = "disable"
	BackEndServicePort = ":8080"
)

func main() {
	// (DSN) for PostgreSQL connection
	dsn := "host=" + DBHost + " user=" + DBUser + " password=" + DBPassword + " dbname=" + DBName + " port=" + DBPort + " sslmode=" + SSlMode

	// Database connection setup
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate schema
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
	log.Printf("%sStarting server on port %s...", color.GreenString("INFO: "), BackEndServicePort)

	err = http.ListenAndServe(BackEndServicePort, r)
	if err != nil {
		log.Fatalf("%sServer failed to start: %v", color.RedString("ERROR: "), err)
	}
}
