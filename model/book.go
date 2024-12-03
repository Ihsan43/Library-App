package model

import "time"

type Book struct {
	ID              string    `gorm:"primary_key;type:uuid;"`
	Title           string    `gorm:"size:255;not null;unique"`
	Writer          string    `gorm:"size:255;not null"`
	PublicationYear int       `gorm:"size:255;not null"`
	Category        string    `gorm:"size:255;not null"`
	Publisher       string    `gorm:"size:255;not null"`
	Description     string    `gorm:"type:text;not null"`
	NumberOfPages   int       `gorm:"size:255;not null"`
	Price           int       `gorm:"size:255;not null"`
	Stock           string    `gorm:"size:255;not null"`
	CoverImg        string
	StockAvailable  int       `gorm:"size:255;not null"`
	CreatedAt       time.Time `gorm:"default:current_timestamp"`
	UpdatedAt       time.Time `gorm:"default:null"`
}
