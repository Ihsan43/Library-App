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
	Register(payload dto.UserRequestDto) (dto.UserResponseDto, error)
	Login(username, password string) (dto.AuthResponseDto, error)
}

type userServi struct {
	us UserService
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
func (s *userServi) Register(payload dto.UserRequestDto) (dto.UserResponseDto, error) {

	exist, err := s.us.CheckEmailOrUsername(payload.Email, payload.Username)
	if err != nil {
		return dto.UserResponseDto{}, err
	}

	if exist {
		return dto.UserResponseDto{}, errors.New("username or email already exists")
	}

	if payload.Email == "" && payload.PhoneNumber == "" && payload.Username == "" && payload.Password == "" {
		return dto.UserResponseDto{}, err
	}

	hashPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return dto.UserResponseDto{}, fmt.Errorf("failed to hash password: %w", err)
	}

	if payload.Role != "admin" && payload.Role != "user" && payload.Role != "employee" {
		return dto.UserResponseDto{}, errors.New("Invalid Role")
	}

	user := model.User{
		Name:        payload.Name,
		Email:       payload.Email,
		PhoneNumber: payload.PhoneNumber,
		Username:    payload.Username,
		Password:    hashPassword,
		Role:        payload.Role,
	}

	newUser, err := s.us.CreateUser(user)
	if err != nil {
		return dto.UserResponseDto{}, fmt.Errorf("failed to create account: %w", err)
	}

	return dto.NewUserResponseDto(newUser), nil
}

func NewAuthService(userService UserService) AuthService {
	return &userServi{
		us: userService,
	}
}
