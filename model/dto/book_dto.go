package dto

import (
	"library_app/model"
	"time"
)

type BookRequestDto struct {
	Title           string `json:"title"`
	Writer          string `json:"writer"`
	PublicationYear int    `json:"publication_year"`
	Category        string `json:"category"`
	Publisher       string `json:"publisher"`
	Description     string `json:"description"`
	CoverImg        string
	Price           int    `json:"price"`
	NumberOfPages   int    `json:"number_of_pages"`
	Stock           string `json:"stock"`
	StockAvailable  int    `json:"stock_available"`
}

type BookResponseDto struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Writer          string    `json:"writer"`
	PublicationYear int       `json:"publication_year"`
	Category        string    `json:"category"`
	Publisher       string    `json:"publisher"`
	Description     string    `json:"description"`
	NumberOfPages   int       `json:"number_of_pages"`
	Stock           string    `json:"stock"`
	StockAvailable  int       `json:"stock_available"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func NewBookResponseDto(book model.Book) BookResponseDto {
	return BookResponseDto{
		ID:              book.ID,
		Title:           book.Title,
		Writer:          book.Writer,
		PublicationYear: book.PublicationYear,
		Category:        book.Category,
		Publisher:       book.Publisher,
		Description:     book.Description,
		NumberOfPages:   book.NumberOfPages,
		Stock:           book.Stock,
		StockAvailable:  book.StockAvailable,
		CreatedAt:       book.CreatedAt,
		UpdatedAt:       book.UpdatedAt,
	}
}
