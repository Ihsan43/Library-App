package repository

import (
	"library_app/model"
	"library_app/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(payload model.User) (model.User, error)
	Get(id string) (model.User, error)
	GetByUsername(username string) (model.User, error)
	Gets() ([]model.User, error)
	Update(id string, payload model.User) (model.User, error)
	IsEmailOrUsernameExist(email, username string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func (a *userRepository) IsEmailOrUsernameExist(email string, username string) (bool, error) {
	var count int64

	if err := a.db.Model(&model.User{}).Where("email = ? OR username = ?", email, username).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (a *userRepository) Create(payload model.User) (model.User, error) {
	payload.ID = utils.GenerateUuid()

	if err := a.db.Create(&payload).Error; err != nil {
		return model.User{}, err
	}

	return payload, nil
}

func (a *userRepository) Get(id string) (model.User, error) {
	var account model.User
	if err := a.db.First(&account, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}
	return account, nil
}

func (a *userRepository) GetByUsername(username string) (model.User, error) {
	var account model.User
	if err := a.db.First(&account, "username = ?", username).Error; err != nil {
		return model.User{}, err
	}
	return account, nil
}

func (a *userRepository) Gets() ([]model.User, error) {
	var accounts []model.User
	if err := a.db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (a *userRepository) Update(id string, payload model.User) (model.User, error) {
	var account model.User
	if err := a.db.First(&account, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}

	if err := a.db.Model(&account).Updates(payload).Error; err != nil {
		return model.User{}, err
	}

	return account, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
