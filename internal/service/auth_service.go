package service

import (
	"errors"
	"library_app/model"
	"library_app/utils"
	"time"
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
		return model.Account{}, err
	}

	if payload.Role != "admin" && payload.Role != "user" && payload.Role != "employee" {
		return model.Account{}, errors.New("Invalid Role")
	}

	payload = model.Account{
		Password:  hashPassword,
		CreatedAt: time.Now(),
	}

	return payload, nil
}

func NewAuthService(accountService AccountService) AuthService {
	return &accountSevi{
		AccountService: accountService,
	}
}
