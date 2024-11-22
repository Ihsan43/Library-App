package service

import (
	"errors"
	"library_app/internal/repository"
	"library_app/model"
	"library_app/utils"
)

type UserService interface {
	CreateUser(payload model.User) (model.User, error)
	CheckEmailOrUsername(email, username string) (bool, error)
	FindByUsernamePassword(username string, password string) (model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// FindUsername implements AccountService.
func (s userService) FindByUsernamePassword(username string, password string) (model.User, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return model.User{}, errors.New("invalid username or password")
	}

	if err := utils.VerifyPassword(user.Password, password); err != nil {
		return model.User{}, errors.New("invalid username or password")
	}

	return user, nil
}

// CheckEmailOrUsername implements AccountService.
func (s *userService) CheckEmailOrUsername(email string, username string) (bool, error) {
	return s.userRepo.IsEmailOrUsernameExist(email, username)
}

// CreateAccount implements AccountService.
func (s *userService) CreateUser(payload model.User) (model.User, error) {
	return s.userRepo.Create(payload)
}

func NewAccountService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
