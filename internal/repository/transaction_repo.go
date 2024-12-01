package repository

import (
	"errors"
	"fmt"
	"library_app/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FetchTransactionHistoriesByUser(userID string) ([]model.TransactionHistory, error)
}

type transactionRepository struct {
	db *gorm.DB
}

// GetTransactionHistoryByPaymentID implements TransactionRepository.
func (r *transactionRepository) FetchTransactionHistoriesByUser(userID string) ([]model.TransactionHistory, error) {
	var histories []model.TransactionHistory

	// Query ke database untuk mengambil semua transaksi berdasarkan user ID
	if err := r.db.Preload("Payment").Preload("Order").
		Where("order_id IN (SELECT id FROM orders WHERE user_id = ?)", userID).Find(&histories).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch transaction histories for user: %v", err)
	}
	if len(histories) == 0 {
		return nil, errors.New("no transaction histories found for the user")
	}

	return histories, nil
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
