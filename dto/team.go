package dto

import "time"

type TeamData struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null"`
	Division  string    `json:"division" gorm:"not null"`
	Year      string    `json:"year" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null"`
	Category  string    `json:"category" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TeamRequest struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Division string `json:"division"`
	Year     string `json:"year"`
	Status   string `json:"status"`
	Category string `json:"category"`
}

type TeamResponse struct {
	StatusCode int        `json:"status"`
	Message    string     `json:"message"`
	Data       []TeamData `json:"data,omitempty"`
}
