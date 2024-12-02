package model

import "time"

type User struct {
	ID          string    `gorm:"primary_key;type:uuid;"`
	Name        string    `gorm:"size:255;not null"`
	Email       string    `gorm:"unique;size:100;not null"`
	PhoneNumber string    `gorm:"unique;size:100;not null"`
	Username    string    `gorm:"unique;size:100;not null"`
	Password    string    `gorm:"size:255"`
	Role        string    `gorm:"size:50;not null"`
	IsVerified  bool      `gorm:"default:false"`
	Status      string    `gorm:"default:active"`
	AvatarURL   string
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"default:null"`
}


