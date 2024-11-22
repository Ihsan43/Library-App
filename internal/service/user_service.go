package service

import (
	"library_app/internal/repository"
	"library_app/model"
)

type UserService interface {
	CreateUser(payload model.User) (model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// CreateAccount implements AccountService.
func (s *userService) CreateUser(payload model.User) (model.User, error) {
	return s.userRepo.Create(payload)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
