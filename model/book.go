package model

import "time"

type Book struct {
	ID              string    `gorm:"primary_key;type:uuid;" json:"id"`
	Title           string    `gorm:"size:255;not null" json:"title"`
	Writer          string    `gorm:"size:255;not null" json:"writer"`
	PublicationYear int       `gorm:"size:255;not null" json:"publication_year"`
	Category        string    `gorm:"size:255;not null" json:"category"`
	Publisher       string    `gorm:"size:255;not null" json:"publisher"`
	Description     string    `gorm:"type:text;not null" json:"description"`
	NumberOfPages   int       `gorm:"size:255;not null" json:"number_of_pages"`
	Stock           int       `gorm:"size:255;not null" json:"stock"`
	StockAvailable  int       `gorm:"size:255;not null" json:"stoak_available"`
	CreatedAt       time.Time `gorm:"default:current_timestamp" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"default:null" json:"updatedAt"`
}
