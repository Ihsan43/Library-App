package service

import (
	"errors"
	"library_app/internal/repository"
	"library_app/model"
	"library_app/model/dto"
	"library_app/utils"
	"library_app/utils/common"
)

type UserService interface {
	CreateUser(payload model.User) (model.User, error)
	CheckEmailOrUsername(email, username string) (bool, error)
	FindByUsernamePassword(username string, password string) (model.User, error)
	FindUserById(id string) (model.User, error)
	UpdatedUser(id string, ppayload dto.UserDto) (model.User, error)
	FindUsers(page, limit int) ([]model.User, int64, error)
	DeleteUserById(id string) (model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// DeleteUserById implements UserService.
func (s *userService) DeleteUserById(id string) (model.User, error) {
	return s.userRepo.DeleteUser(id)
}

// FindUsers implements UserService.
func (s *userService) FindUsers(page int, limit int) ([]model.User, int64, error) {
	paginator := common.NewPaginator(page, limit)
	return s.userRepo.GetUsers(paginator)
}

func (s *userService) UpdatedUser(id string, payload dto.UserDto) (model.User, error) {

	user, err := s.userRepo.GetUser(id)
	if err != nil {
		return model.User{}, errors.New("user not found")
	}

	user = model.User{
		Name:        payload.Name,
		Username:    payload.Username,
		PhoneNumber: payload.PhoneNumber,
		AvatarURL:   payload.AvatarURL,
	}

	newUser, err := s.userRepo.UpdateUser(id, user)
	if err != nil {
		return model.User{}, errors.New(err.Error())
	}

	return newUser, nil
}

func (s *userService) FindUserById(id string) (model.User, error) {
	return s.userRepo.GetUser(id)
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
	return s.userRepo.CreateUser(payload)
}

func NewAccountService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
