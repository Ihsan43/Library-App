package repository

import (
	"library_app/model"
	"library_app/utils"
	"library_app/utils/common"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(payload model.User) (model.User, error)
	GetUser(id string) (model.User, error)
	GetByUsername(username string) (model.User, error)
	GetUsers(paginator *common.Paginator) ([]model.User, int64, error)
	UpdateUser(id string, payload model.User) (model.User, error)
	IsEmailOrUsernameExist(email, username string) (bool, error)
	DeleteUser(id string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// DeleteUser implements UserRepository.
func (a *userRepository) DeleteUser(id string) (model.User, error) {
	var user model.User

	if err := a.db.Model(&model.User{}).First(&user, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}

	if err := a.db.Model(&model.User{}).Delete(&user, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (a *userRepository) IsEmailOrUsernameExist(email string, username string) (bool, error) {
	var count int64

	if err := a.db.Model(&model.User{}).Where("email = ? OR username = ?", email, username).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (a *userRepository) CreateUser(payload model.User) (model.User, error) {
	payload.ID = utils.GenerateUuid()

	if err := a.db.Create(&payload).Error; err != nil {
		return model.User{}, err
	}

	return payload, nil
}

func (a *userRepository) GetUser(id string) (model.User, error) {
	var user model.User
	if err := a.db.First(&user, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a *userRepository) GetByUsername(username string) (model.User, error) {
	var user model.User
	if err := a.db.First(&user, "username = ?", username).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a *userRepository) GetUsers(paginator *common.Paginator) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	// Hitung total item
	a.db.Model(&model.User{}).Count(&total)

	// Terapkan pagination dan ambil data
	err := paginator.ApplyPagination(a.db).Find(&users).Error
	return users, total, err
}

func (a *userRepository) UpdateUser(id string, payload model.User) (model.User, error) {
	var user model.User
	if err := a.db.First(&user, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}

	if err := a.db.Model(&user).Updates(payload).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
