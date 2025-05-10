package model

import "time"

type Team struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null"`
	Division  string    `json:"division" gorm:"not null"`
	Year      string    `json:"year" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
}

func (m *Team) TableName() string {
	return "team"
}
