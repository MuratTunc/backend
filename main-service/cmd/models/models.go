package models

// ImageData represents the image metadata to be stored in the database
type ImageData struct {
	ID           uint   `gorm:"primaryKey"`   // Auto-incrementing ID field
	Title        string `json:"title"`        // Title of the image or case study
	Text         string `json:"text"`         // Description of the image or case study
	ImageURL     string `json:"imageUrl"`     // URL of the image
	CreationTime string `json:"creationTime"` // Time when the image was created in FireBase Cloud
}
