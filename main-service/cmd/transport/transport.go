package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"main-service/cmd/models"
	"main-service/cmd/service"

	"github.com/go-kit/kit/endpoint"
)

// SaveImageDataRequest defines the incoming request structure.
type SaveImageDataRequest struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	ImageURL     string `json:"imageUrl"`
	CreationTime string `json:"creationTime"`
}

// SaveImageDataResponse defines the outgoing response structure.
type SaveImageDataResponse struct {
	Message string `json:"message"`
	Err     string `json:"error,omitempty"` // Optional error field
}

// MakeSaveImageDataEndpoint creates the endpoint for saving image metadata.
func MakeSaveImageDataEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SaveImageDataRequest)

		// Convert the request into a models.ImageData struct
		imageData := models.ImageData{
			Title:        req.Title,
			Description:  req.Description,
			ImageURL:     req.ImageURL,
			CreationTime: req.CreationTime,
		}

		// Call the service method
		err := svc.SaveImageData(ctx, imageData)
		if err != nil {
			return SaveImageDataResponse{Message: "", Err: err.Error()}, nil
		}

		// If successful, return a success message
		return SaveImageDataResponse{Message: "Image metadata saved successfully"}, nil
	}
}

// DecodeSaveImageDataRequest decodes incoming JSON requests.
func DecodeSaveImageDataRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request SaveImageDataRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("Failed to decode request: %v", err) // Log any errors that occur during decoding
		return nil, err
	}
	log.Printf("Decoded request: %+v", request) // Log the request to see if it's correctly decoded
	return request, nil
}

// EncodeResponse encodes responses to JSON format.
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
