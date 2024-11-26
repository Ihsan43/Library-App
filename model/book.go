package model

import "time"

type Book struct {
	ID              string `gorm:"primary_key;type:uuid;" json:"id"`
	Title           string `gorm:"size:255;not null;unique" json:"title"`
	Writer          string `gorm:"size:255;not null" json:"writer"`
	PublicationYear int    `gorm:"size:255;not null" json:"publication_year"`
	Category        string `gorm:"size:255;not null" json:"category"`
	Publisher       string `gorm:"size:255;not null" json:"publisher"`
	Description     string `gorm:"type:text;not null" json:"description"`
	NumberOfPages   int    `gorm:"size:255;not null" json:"number_of_pages"`
	Price           int    `gorm:"size:255;not null" json:"price"`
	Stock           string `gorm:"default:stock available;size:255;not null" json:"stock"`
	CoverImg        string
	StockAvailable  int       `gorm:"size:255;not null" json:"stock_available"`
	CreatedAt       time.Time `gorm:"default:current_timestamp" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"default:null" json:"updatedAt"`
}
