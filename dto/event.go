package dto

import "time"

type EventData struct {
	ID               uint64    `json:"id" gorm:"primaryKey"`
	EventName        string    `json:"event_name" gorm:"not null"`
	EventDate        string    `json:"event_date" gorm:"not null"`
	EventTime        string    `json:"event_time" gorm:"not null"`
	EventLocation    string    `json:"event_location" gorm:"not null"`
	EventDescription string    `json:"event_description"`
	EventOrganizer   string    `json:"event_organizer" gorm:"not null"`
	EventStatus      string    `json:"event_status" gorm:"not null"`
	EventImageURL    string    `json:"event_image_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type EventRequest struct {
	EventName        string `json:"event_name"`
	EventDate        string `json:"event_date"`
	EventTime        string `json:"event_time"`
	EventLocation    string `json:"event_location"`
	EventDescription string `json:"event_description"`
	EventOrganizer   string `json:"event_organizer"`
	EventStatus      string `json:"event_status"`
}

type EventResponse struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"message"`
	Data       []EventData `json:"data,omitempty"`
}
