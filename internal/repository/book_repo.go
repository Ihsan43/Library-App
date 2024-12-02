package repository

import (
	"library_app/model"
	"library_app/utils"
	"library_app/utils/common"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(payload model.Book) (model.Book, error)
	GetBook(id string) (model.Book, error)
	Gebooks(paginator *common.Paginator) ([]model.Book, int64, error)
	UpdateBook(id string, payload model.Book) (model.Book, error)
	DeleteBook(id string) (model.Book, error)
}

type bookRepository struct {
	db gorm.DB
}

// UpdateBook implements BookRepository.
func (r *bookRepository) UpdateBook(id string, payload model.Book) (model.Book, error) {
	var book model.Book
	if err := r.db.First(&book, "id = ?", id).Error; err != nil {
		return model.Book{}, err
	}

	if err := r.db.Model(&book).Updates(payload).Error; err != nil {
		return model.Book{}, err
	}

	return book, nil
}

// DeleteBook implements BookRepository.
func (r *bookRepository) DeleteBook(id string) (model.Book, error) {
	var book model.Book

	if err := r.db.Model(&book).First(&book, "id = ?", id).Delete(&book, "id = ?", id).Error; err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (r *bookRepository) Gebooks(paginator *common.Paginator) ([]model.Book, int64, error) {
	var books []model.Book
	var total int64

	r.db.Model(&model.Book{}).Count(&total)

	err := paginator.ApplyPagination(&r.db).Find(&books).Error
	return books, total, err
}

// GetBook implements BookRepository.
func (r *bookRepository) GetBook(id string) (model.Book, error) {
	var book model.Book

	if err := r.db.Model(&book).First(&book, "id = ?", id).Error; err != nil {
		return model.Book{}, err
	}

	return book, nil
}

// CreateBook implements BookRepository.
func (r *bookRepository) CreateBook(payload model.Book) (model.Book, error) {

	payload.ID = utils.GenerateUuid()

	if err := r.db.Model(&model.Book{}).Create(&payload).Error; err != nil {
		return model.Book{}, err
	}

	return payload, nil
}

func NewBookRepository(db gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}
