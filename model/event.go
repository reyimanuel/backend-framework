package model

type Event struct {
	ID               uint64 `json:"id" gorm:"primaryKey"`
	EventName        string `json:"event_name" gorm:"not null"`
	EventDate        string `json:"event_date" gorm:"not null"`
	EventTime        string `json:"event_time" gorm:"not null"`
	EventLocation    string `json:"event_location" gorm:"not null"`
	EventDescription string `json:"event_description" gorm:"not null"`
	EventOrganizer   string `json:"event_organizer" gorm:"not null"`
	EventStatus      string `json:"event_status" gorm:"not null"`
	CreatedAt        string `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	UpdatedAt        string `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
}

func (m *Event) TableName() string {
	return "event"
}
