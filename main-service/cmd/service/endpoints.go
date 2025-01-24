package service

import (
	"context"
	"errors"
	"main-service/cmd/models"
)

// Define the error for an invalid request
var ErrInvalidRequest = errors.New("invalid request")

// SaveImageDataRequest represents the input data for the SaveImageData endpoint.
type SaveImageDataRequest struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	ImageURL     string `json:"imageUrl"`
	CreationTime string `json:"creationTime"`
}

// SaveImageDataResponse represents the output data from the SaveImageData endpoint.
type SaveImageDataResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

// MakeSaveImageDataEndpoint creates the endpoint for saving image metadata.
func MakeSaveImageDataEndpoint(svc Service) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// Convert the incoming request to SaveImageDataRequest
		req, ok := request.(SaveImageDataRequest)
		if !ok {
			return nil, ErrInvalidRequest
		}
		// Mapping to models.ImageData struct
		data := models.ImageData{
			Title:        req.Title,
			Description:  req.Description,
			ImageURL:     req.ImageURL,
			CreationTime: req.CreationTime,
		}
		// Call the service layer to save the image metadata
		err := svc.SaveImageData(ctx, data)
		if err != nil {
			return SaveImageDataResponse{Error: err.Error()}, nil
		}
		// Return a success message in the response
		return SaveImageDataResponse{Message: "Image metadata saved successfully"}, nil
	}
}
