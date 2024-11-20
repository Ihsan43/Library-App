package service

import (
	"errors"
	"library_app/internal/repository"
	"library_app/model"
	"library_app/utils"
)

type AccountService interface {
	CreateAccount(payload model.Account) (model.Account, error)
	CheckEmailOrUsername(email, username string) (bool, error)
	FindByUsernamePassword(username string, password string) (model.Account, error)
}

type accountService struct {
	accountRepo repository.AccountRepository
}

// FindUsername implements AccountService.
func (s accountService) FindByUsernamePassword(username string, password string) (model.Account, error) {
	account, err := s.accountRepo.GetByUsername(username)
	if err != nil {
		return model.Account{}, errors.New("invalid usernamesada or password")
	}

	if err := utils.VerifyPassword(account.Password, password); err != nil {
		return model.Account{}, errors.New("invalid username or password")
	}

	return account,nil
}

// CheckEmailOrUsername implements AccountService.
func (s *accountService) CheckEmailOrUsername(email string, username string) (bool, error) {
	return s.accountRepo.IsEmailOrUsernameExist(email, username)
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
