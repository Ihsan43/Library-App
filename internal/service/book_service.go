package service

import (
	"errors"
	"fmt"
	"library_app/internal/repository"
	"library_app/model"
	"library_app/model/dto"
	"library_app/utils/common"
)

type BookService interface {
	CreateBook(payload dto.BookRequestDto) (dto.BookResponseDto, error)
	GetBookById(id string) (dto.BookResponseDto, error)
	FindBooks(page, limit int) ([]dto.BookResponseDto, int64, error)
	DeleteBookById(id string) (model.Book, error)
	UpdateBookById(id string, payload dto.BookRequestDto) (dto.BookResponseDto, error)
}

type bookService struct {
	bookRepo repository.BookRepository
}

// UpdateBookById implements BookService.
func (s *bookService) UpdateBookById(id string, payload dto.BookRequestDto) (dto.BookResponseDto, error) {
	if payload.PublicationYear < 1000 || payload.PublicationYear > 9999 {
		return dto.BookResponseDto{}, errors.New("invalid publication year: must be a 4-digit number")
	}

	var stockStatus string
	if payload.StockAvailable == 0 {
		stockStatus = "Stock is empty"
	} else {
		stockStatus = "Stock available"
	}

	book := model.Book{
		Title:           payload.Title,
		Writer:          payload.Writer,
		PublicationYear: payload.PublicationYear,
		Category:        payload.Category,
		Publisher:       payload.Publisher,
		Description:     payload.Description,
		NumberOfPages:   payload.NumberOfPages,
		Price:           payload.Price,
		StockAvailable:  payload.StockAvailable,
		Stock:           stockStatus,
	}

	newBook, err := s.bookRepo.UpdateBook(id, book)
	if err != nil {
		return dto.BookResponseDto{}, fmt.Errorf("failed to update book: %v", err)
	}

	return dto.NewBookResponseDto(newBook), nil

}

// DeleteBookById implements BookService.
func (s *bookService) DeleteBookById(id string) (model.Book, error) {
	return s.bookRepo.DeleteBook(id)
}

// FindBooks implements BookService.
func (s *bookService) FindBooks(page, limit int) ([]dto.BookResponseDto, int64, error) {
	paginator := common.NewPaginator(page, limit)
	books, total, err := s.bookRepo.Gebooks(paginator)
	if err != nil {
		return nil, 0, err
	}

	// Gunakan helper untuk konversi
	var bookDtos []dto.BookResponseDto
	for _, book := range books {
		bookDtos = append(bookDtos, dto.NewBookResponseDto(book))
	}

	return bookDtos, total, nil

}

// GetBookById implements BookService.
func (s *bookService) GetBookById(id string) (dto.BookResponseDto, error) {

	book, err := s.bookRepo.GetBook(id)
	if err != nil {
		return dto.BookResponseDto{}, err
	}

	return dto.NewBookResponseDto(book), nil
}

// CreateBook implements BookService.
func (s *bookService) CreateBook(payload dto.BookRequestDto) (dto.BookResponseDto, error) {

	if payload.PublicationYear < 1000 || payload.PublicationYear > 9999 {
		return dto.BookResponseDto{}, errors.New("invalid publication year: must be a 4-digit number")
	}

	var stockStatus string
	if payload.StockAvailable == 0 {
		stockStatus = "Stock is empty"
	} else {
		stockStatus = "Stock available"
	}

	book := model.Book{
		Title:           payload.Title,
		Writer:          payload.Writer,
		PublicationYear: payload.PublicationYear,
		Category:        payload.Category,
		Publisher:       payload.Publisher,
		Description:     payload.Description,
		NumberOfPages:   payload.NumberOfPages,
		Price:           payload.Price,
		StockAvailable:  payload.StockAvailable,
		Stock:           stockStatus,
	}

	newBook, err := s.bookRepo.CreateBook(book)
	if err != nil {
		return dto.BookResponseDto{}, fmt.Errorf("failed to create book: %v", err)
	}

	return dto.NewBookResponseDto(newBook), nil
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}
