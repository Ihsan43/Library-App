package service

import (
	"library_app/internal/repository"
	"library_app/model"
)

type AccountService interface {
	CreateAccount(payload model.Account) (model.Account, error)
}

type accountService struct {
	accountRepo repository.AccountRepository
}

// CreateAccount implements AccountService.
func (s *accountService) CreateAccount(payload model.Account) (model.Account, error) {
	return s.accountRepo.Create(payload)
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return &accountService{
		accountRepo: accountRepo,
	}
}
