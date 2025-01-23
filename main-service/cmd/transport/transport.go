package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"main-service/cmd/models"
	"main-service/cmd/service"

	"github.com/go-kit/kit/endpoint"
)

// SaveImageDataRequest defines the incoming request structure.
type SaveImageDataRequest struct {
	ImageData models.ImageData `json:"imageData"` // Metadata
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

		// Call the service method
		err := svc.SaveImageData(ctx, req.ImageData)
		if err != nil {
			return SaveImageDataResponse{Message: "", Err: err.Error()}, nil
		}

		// If successful, return a success message
		return SaveImageDataResponse{Message: "Image metadata saved successfully"}, nil
	}
}

// DecodeSaveImageDataRequest decodes incoming JSON requests.
func DecodeSaveImageDataRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req SaveImageDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// EncodeResponse encodes responses to JSON format.
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
