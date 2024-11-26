package model

import "time"

type Address struct {
	ID          string    `gorm:"primary_key;type:uuid;" json:"id"`
	UserID      string    `gorm:"type:uuid;not null" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	PhoneNumber string    `gorm:"size:255;not null" json:"phone_number"`
	Street      string    `gorm:"size:255;not null" json:"street"`
	City        string    `gorm:"size:255;not null" json:"city"`
	PostalCode  string    `gorm:"size:20;not null" json:"postal_code"`
	State       string    `gorm:"size:255;not null" json:"state"`
	Country     string    `gorm:"size:255;not null" json:"country"`
	CreatedAt   time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:null" json:"updated_at"`
}
