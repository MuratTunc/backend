package main

import (
	"log"
	"net/http"

	kitHttp "github.com/go-kit/kit/transport/http" // Alias Go-Kit HTTP transport
	"github.com/gorilla/mux"                       // Gorilla Mux router
	"github.com/rs/cors"                           // CORS package
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
	DBPassword         = "depixen-pass"
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
		red := color.New(color.FgRed).SprintFunc()
		log.Fatalf(red("Failed to connect to the database: %v"), err)
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

	// Enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000", // Allow requests from localhost (your React frontend)
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Cache preflight response for 5 minutes
	}).Handler(r)

	// Start the HTTP server
	log.Printf("%sStarting server on port %s...", color.GreenString("INFO: "), BackEndServicePort)

	err = http.ListenAndServe(BackEndServicePort, corsHandler)
	if err != nil {
		log.Fatalf("%sServer failed to start: %v", color.RedString("ERROR: "), err)
	}
}
