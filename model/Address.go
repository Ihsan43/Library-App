package model

import "time"

type Address struct {
	ID          string    `gorm:"primary_key;type:uuid;"`
	UserID      string    `gorm:"type:uuid;not null"`
	User        User      `gorm:"foreignKey:UserID;references:ID"`
	PhoneNumber string    `gorm:"size:255;not null"`
	Street      string    `gorm:"size:255;not null"`
	City        string    `gorm:"size:255;not null"`
	PostalCode  string    `gorm:"size:20;not null"`
	State       string    `gorm:"size:255;not null"`
	Country     string    `gorm:"size:255;not null"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"default:null"`
}
