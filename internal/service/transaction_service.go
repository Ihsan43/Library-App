package service

import (
	"errors"
	"library_app/internal/repository"
	"library_app/model"
)

type TransactionService interface {
	GetTransactionHistoriesByUser(userID string) ([]model.TransactionHistory, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
	}
}

func (s *transactionService) GetTransactionHistoriesByUser(userID string) ([]model.TransactionHistory, error) {
	// Validasi jika userID kosong
	if userID == "" {
		return nil, errors.New("userID cannot be empty")
	}

	histories, err := s.transactionRepository.FetchTransactionHistoriesByUser(userID)
	if err != nil {
		return nil, err // Jangan tambahkan pesan baru
	}

	return histories, nil
}
