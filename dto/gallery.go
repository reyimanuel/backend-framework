package dto

import "time"

type GalleryData struct {
	ID          uint64    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url" gorm:"not null"`
	Category   string    `json:"category" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GalleryRequest struct {
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	Category    string `json:"category" gorm:"not null"`
	// ImageURL    string `json:"image_url" gorm:"not null"`
}

type GalleryResponse struct {
	StatusCode int           `json:"status"`
	Message    string        `json:"message"`
	Data       []GalleryData `json:"data,omitempty"`
}
