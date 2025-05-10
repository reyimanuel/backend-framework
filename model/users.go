package model

type User struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	Email     string `json:"email" gorm:"unique;not null"`
	Username  string `json:"username" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	NIM       int    `json:"nim" gorm:"unique;not null"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}

func (u *User) TableName() string {
	return "users"
}
