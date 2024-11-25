package service

import (
	"errors"
	"library_app/internal/repository"
	"library_app/model"
	"library_app/utils/common"
)

type BookService interface {
	CreateBook(payload model.Book) (model.Book, error)
	GetBookById(id string) (model.Book, error)
	FindBooks(page, limit int) ([]model.Book, int64, error)
	DeleteBookById(id string) (model.Book, error)
}

type bookService struct {
	bookRepo repository.BookRepository
}

// DeleteBookById implements BookService.
func (s *bookService) DeleteBookById(id string) (model.Book, error) {
	return s.bookRepo.DeleteBook(id)
}

// FindBooks implements BookService.
func (s *bookService) FindBooks(page int, limit int) ([]model.Book, int64, error) {
	paginator := common.NewPaginator(page, limit)
	return s.bookRepo.Gebooks(paginator)
}

// GetBookById implements BookService.
func (s *bookService) GetBookById(id string) (model.Book, error) {
	return s.bookRepo.GetBook(id)
}

// CreateBook implements BookService.
func (s *bookService) CreateBook(payload model.Book) (model.Book, error) {

	if payload.PublicationYear < 1000 || payload.PublicationYear > 9999 {
		return model.Book{}, errors.New("invalid format year: must be a 4-digit number")
	}

	if payload.StockAvailable == 0 {
		payload.Stock = "Stock is empty"
	}

	book, err := s.bookRepo.CreateBook(payload)
	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}
