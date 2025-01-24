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

	"github.com/fatih/color" // color package
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

func setupCORS(r *mux.Router) http.Handler {
	// Enable CORS
	return cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000", // Allow requests from localhost our React frontend web-app
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Cache preflight response for 5 minutes
	}).Handler(r)
}

func connectToDB() (*gorm.DB, error) {
	// (DSN) for PostgreSQL connection
	dsn := "host=" + DBHost + " user=" + DBUser + " password=" + DBPassword + " dbname=" + DBName + " port=" + DBPort + " sslmode=" + SSlMode
	// Database connection setup
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.ImageData{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initializeService(db *gorm.DB) service.Service {
	// Initialize service (passing DB connection to the service layer)
	return service.NewService(db)
}

func setupRoutes(svc service.Service) *mux.Router {
	// Set up Gorilla Mux router
	r := mux.NewRouter()
	// Create the endpoint for saving image data
	saveImageDataEndpoint := transport.MakeSaveImageDataEndpoint(svc)

	// Create HTTP handler using Go-Kit
	saveImageDataHandler := kitHttp.NewServer(
		saveImageDataEndpoint,                // The endpoint we created
		transport.DecodeSaveImageDataRequest, // Request decoder
		transport.EncodeResponse,             // Response encoder
	)

	// Define the route for saving image data
	r.Handle("/api/v1/main-service/save-image-data", saveImageDataHandler).Methods("POST")

	return r
}

func startHTTPServer(r *mux.Router) error {
	// Call the setupCORS function to get the CORS handler
	corsHandler := setupCORS(r)
	// Start the HTTP server
	log.Printf("%sStarting server on port %s...", color.GreenString("INFO: "), BackEndServicePort)
	return http.ListenAndServe(BackEndServicePort, corsHandler)
}

func main() {
	// Connect to the database
	db, err := connectToDB()
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		log.Fatalf(red("Failed to connect to the database: %v"), err)
	}
	// Initialize service
	svc := initializeService(db)

	// Set up routes
	r := setupRoutes(svc)

	// Start the HTTP server
	if err := startHTTPServer(r); err != nil {
		log.Fatalf("%sServer failed to start: %v", color.RedString("ERROR: "), err)
	}
}
