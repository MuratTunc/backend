package service

import (
	"context"
	"main-service/cmd/models"

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
	// Insert the ImageData into the database using GORM
	if err := s.DB.WithContext(ctx).Create(&data).Error; err != nil {
		return err
	}
	return nil // Return nil to indicate success
}
