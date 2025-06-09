package model

import "time"

type Event struct {
	ID               uint64    `json:"id" gorm:"primaryKey"`
	EventName        string    `json:"event_name" gorm:"not null"`
	EventDate        time.Time `json:"event_date" gorm:"not null"`
	EventTime        string    `json:"event_time" gorm:"not null"`
	EventLocation    string    `json:"event_location" gorm:"not null"`
	EventDescription string    `json:"event_description" gorm:"not null"`
	EventOrganizer   string    `json:"event_organizer" gorm:"not null"`
	EventStatus      string    `json:"event_status" gorm:"not null"`
	EventCategory    string    `json:"event_category" gorm:"not null"`
	EventImageURL    string    `json:"event_image_url" gorm:"not null"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	UpdatedAt        time.Time `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
}

func (m *Event) TableName() string {
	return "event"
}
