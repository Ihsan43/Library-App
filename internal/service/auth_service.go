package service

import (
	"errors"
	"fmt"
	"library_app/model"
	"library_app/model/dto"
	"library_app/utils"
	"library_app/utils/security"
)

type AuthService interface {
	Register(payload model.User) (model.User, error)
	Login(username, password string) (dto.AuthResponseDto, error)
}

type userServi struct {
	us          UserService
}

// Login implements AuthService.
func (s *userServi) Login(username string, password string) (dto.AuthResponseDto, error) {

	account, err := s.us.FindByUsernamePassword(username, password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	token, err := security.CreateAccessToken(&account)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return token, nil
}

// RegisterAccount implements AuthService.
func (s *userServi) Register(payload model.User) (model.User, error) {

	exist, err := s.us.CheckEmailOrUsername(payload.Email, payload.Username)
	if err != nil {
		return model.User{}, err
	}

	if exist {
		return model.User{}, errors.New("username or email already exists")
	}

	hashPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	if payload.Role != "admin" && payload.Role != "user" && payload.Role != "employee" {
		return model.User{}, errors.New("Invalid Role")
	}

	payload.Password = hashPassword

	account, err := s.us.CreateUser(payload)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to create account: %w", err)
	}

	return account, nil
}

func NewAuthService(userService UserService) AuthService {
	return &userServi{
		us: userService,
	}
}
