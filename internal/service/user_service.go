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
	FindUserById(id string) (dto.UserResponseDto, error) 
	UpdatedUser(id string, payload dto.UserRequestDto) (dto.UserResponseDto, error)
	FindUsers(page int, limit int) ([]dto.UserResponseDto, int64, error)
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
func (s *userService) FindUsers(page int, limit int) ([]dto.UserResponseDto, int64, error) {
	paginator := common.NewPaginator(page, limit)
	users, total, err := s.userRepo.GetUsers(paginator)
	if err != nil {
		return nil, 0, err
	}

	// Gunakan helper untuk konversi
	var userDtos []dto.UserResponseDto
	for _, user := range users {
		userDtos = append(userDtos, dto.NewUserResponseDto(user))
	}

	return userDtos, total, nil
}


func (s *userService) UpdatedUser(id string, payload dto.UserRequestDto) (dto.UserResponseDto, error) {

	user, err := s.userRepo.GetUser(id)
	if err != nil {
		return dto.UserResponseDto{}, errors.New("user not found")
	}

	user = model.User{
		Name:        payload.Name,
		Username:    payload.Username,
		AvatarURL:   payload.AvatarURL,
	}

	newUser, err := s.userRepo.UpdateUser(id, user)
	if err != nil {
		return dto.UserResponseDto{}, errors.New(err.Error())
	}
	
	return dto.NewUserResponseDto(newUser), nil
}

func (s *userService) FindUserById(id string) (dto.UserResponseDto, error) {
	
	user, err := s.userRepo.GetUser(id)
	if err != nil {
		return dto.UserResponseDto{}, errors.New("user not found")
	} 

	return dto.NewUserResponseDto(user), nil

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
