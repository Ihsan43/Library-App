package service

import (
	"errors"
	"library_app/model"
	"library_app/utils"
	"time"
)

type AuthService interface {
	RegisterUser(payload model.User) (model.User, error)
}

type userSevi struct {
	UserService
}

// RegisterAccount implements AuthService.
func (s *userSevi) RegisterUser(payload model.User) (model.User, error) {

	hashPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return model.User{}, err
	}

	if payload.Role != "admin" && payload.Role != "user" && payload.Role != "employee" {
		return model.User{}, errors.New("Invalid Role")
	}

	payload = model.User{
		Password:  hashPassword,
		CreatedAt: time.Now(),
	}

	return payload, nil
}

func NewAuthService(userService UserService) AuthService {
	return &userSevi{
		UserService: userService,
	}
}
