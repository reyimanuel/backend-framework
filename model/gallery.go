package model

import "time"

type Gallery struct {
	ID          uint64    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	ImageURL    string    `json:"image_url" gorm:"not null"`
	Category    string    `json:"category" gorm:"not null"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
}

func (m *Gallery) TableName() string {
	return "gallery"
}
