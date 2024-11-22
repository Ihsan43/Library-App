package model

import "time"

type User struct {
	ID        string    `gorm:"primary_key;type:uuid;" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Email     string    `gorm:"unique;size:100;not null" json:"email"`
	Username  string    `gorm:"unique;size:100;not null" json:"username"`
	Password  string    `gorm:"size:255" json:"password,omitempty"` 
	Role      string    `gorm:"size:50;not null" json:"role"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:current_timestamp;autoUpdateTime" json:"updatedAt"`
}

