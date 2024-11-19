package service

import (
	"errors"
	"fmt"
	"library_app/model"
	"library_app/utils"
)

type AuthService interface {
	RegisterAccount(payload model.Account) (model.Account, error)
}

type accountSevi struct {
	AccountService
}

// RegisterAccount implements AuthService.
func (s *accountSevi) RegisterAccount(payload model.Account) (model.Account, error) {

	hashPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return model.Account{}, fmt.Errorf("failed to hash password: %w", err)
	}

	if payload.Role != "admin" && payload.Role != "user" && payload.Role != "employee" {
		return model.Account{}, errors.New("Invalid Role")
	}

	payload.Password = hashPassword

	account, err := s.CreateAccount(payload)
	if err != nil {
		return model.Account{}, fmt.Errorf("failed to create account: %w", err)
	}

	return account, nil
}

func NewAuthService(accountService AccountService) AuthService {
	return &accountSevi{
		AccountService: accountService,
	}
}
