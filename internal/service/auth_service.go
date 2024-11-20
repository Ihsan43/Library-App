package service

import (
	"errors"
	"fmt"
	"library_app/model"
	"library_app/utils"
)

type AuthService interface {
	RegisterAccount(payload model.Account) (model.Account, error)
	Login(username, password string) (model.Account, error)
}

type accountSevi struct {
	AccountService
}

// Login implements AuthService.
func (s *accountSevi) Login(username string, password string) (model.Account, error) {
	account, err := s.AccountService.FindByUsernamePassword(username, password)
	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}

// RegisterAccount implements AuthService.
func (s *accountSevi) RegisterAccount(payload model.Account) (model.Account, error) {

	exist, err := s.AccountService.CheckEmailOrUsername(payload.Email, payload.Username)
	if err != nil {
		return model.Account{}, err
	}

	if exist {
		return model.Account{}, errors.New("username or email already exists")
	}

	hashPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return model.Account{}, fmt.Errorf("failed to hash password: %w", err)
	}

	if payload.Role != "admin" && payload.Role != "user" && payload.Role != "employee" {
		return model.Account{}, errors.New("Invalid Role")
	}

	payload.Password = hashPassword

	account, err := s.AccountService.CreateAccount(payload)
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
