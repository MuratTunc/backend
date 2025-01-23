package service

import (
	"context"
	"log"
	"main-service/cmd/models"

	"github.com/fatih/color" // Import color package for colored logs
	"gorm.io/gorm"
)

// Service defines the service interface.
type Service interface {
	SaveImageData(ctx context.Context, data models.ImageData) error
}

// ImageService implements the Service interface.
type ImageService struct {
	DB *gorm.DB
}

// NewService creates a new ImageService with a DB connection.
func NewService(db *gorm.DB) Service {
	return &ImageService{DB: db}
}

// SaveImageData saves metadata (ImageData) into the database.
func (s *ImageService) SaveImageData(ctx context.Context, data models.ImageData) error {

	log.Printf("Received data: %+v", data) // Log the received data

	// Insert the ImageData into the database
	err := s.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		log.Printf(red("Failed to save data to the database: %v"), err)

		// Return the error for further handling
		return err
	}

	green := color.New(color.FgGreen).SprintFunc()
	log.Printf(green("Successfully saved data to the database: %+v"), data)

	return nil // Success
}
