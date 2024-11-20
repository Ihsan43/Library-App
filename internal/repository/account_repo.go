package repository

import (
	"library_app/model"
	"library_app/utils"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(payload model.Account) (model.Account, error)
	Get(id string) (model.Account, error)
	GetByUsername(username string) (model.Account, error)
	Gets() ([]model.Account, error)
	Update(id string, payload model.Account) (model.Account, error)
	IsEmailOrUsernameExist(email, username string) (bool, error)

}

type accountRepository struct {
	db *gorm.DB
}

// IsEmailOrUsernameExist implements AccountRepository.
func (a *accountRepository) IsEmailOrUsernameExist(email string, username string) (bool, error) {
	var count int64

	if err := a.db.Model(&model.Account{}).Where("email = ? OR username = ?", email, username).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

// Create implements AccountRepository.
func (a *accountRepository) Create(payload model.Account) (model.Account, error) {
	payload.ID = utils.GenerateUuid()

	if err := a.db.Create(&payload).Error; err != nil {
		return model.Account{}, err
	}

	return payload, nil
}

func (a *accountRepository) Get(id string) (model.Account, error) {
	var account model.Account
	if err := a.db.First(&account, "id = ?", id).Error; err != nil {
		return model.Account{}, err
	}
	return account, nil
}

// GetByUsername implements AccountRepository.
func (a *accountRepository) GetByUsername(username string) (model.Account, error) {
	var account model.Account
	if err := a.db.First(&account, "username = ?", username).Error; err != nil {
		return model.Account{}, err
	}
	return account, nil
}

// Gets implements AccountRepository.
func (a *accountRepository) Gets() ([]model.Account, error) {
	var accounts []model.Account
	if err := a.db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

// Update implements AccountRepository.
func (a *accountRepository) Update(id string, payload model.Account) (model.Account, error) {
	var account model.Account
	if err := a.db.First(&account, "id = ?", id).Error; err != nil {
		return model.Account{}, err
	}

	if err := a.db.Model(&account).Updates(payload).Error; err != nil {
		return model.Account{}, err
	}

	return account, nil
}

// NewAccountRepository menginisialisasi dan mengembalikan AccountRepository.
func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}
