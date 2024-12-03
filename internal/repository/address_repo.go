package repository

import (
	"errors"
	"library_app/model"
	"library_app/utils"

	"gorm.io/gorm"
)

type AddressRepository interface {
	CreateAddress(address model.Address) (model.Address, error)
	UpdateAddress(id string, payload model.Address) (model.Address, error)
	DeleteAddress(id string) error
	GetAddress(id string) (model.Address, error)
	FindAddressByUserId(userId string) (model.Address, error)
	CountAddress(userId string) (int64, error)
	FindAddresses(userId string) ([]model.Address, error)
}

type addressRepository struct {
	db *gorm.DB
}

// FindAddresses implements AddressRepository.
func (r *addressRepository) FindAddresses(userId string) ([]model.Address, error) {

	var addresses []model.Address

	if err := r.db.Model(&model.Address{}).Preload("User").Where("user_id = ?", userId).Find(&addresses).Error; err != nil {
		return nil, err
	}

	return addresses, nil
}

// CountAddress implements AddressRepository.
func (r *addressRepository) CountAddress(userId string) (int64, error) {

	var count int64

	if err := r.db.Model(&model.Address{}).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// FindAddressByUserId implements AddressRepository.
func (r *addressRepository) FindAddressByUserId(userId string) (model.Address, error) {

	var address model.Address

	if err := r.db.Model(&address).First(&address, "user_id = ?", userId).Error; err != nil {
		return model.Address{}, err
	}

	return address, nil
}

// DeleteAddress implements AddressRepository.
func (r *addressRepository) DeleteAddress(id string) error {

	var address model.Address

	if err := r.db.Model(&address).First(&address, "id = ?", id).Delete(&address).Error; err != nil {
		return err
	}

	return nil
}

// GetAddress implements AddressRepository.
func (r *addressRepository) GetAddress(id string) (model.Address, error) {

	var address model.Address

	if err := r.db.Preload("User").First(&address, "id = ?", id).Error; err != nil {
		return model.Address{}, err
	}

	return address, nil
}

// UpdateAddress implements AddressRepository.
func (r *addressRepository) UpdateAddress(id string, payload model.Address) (model.Address, error) {
	var address model.Address

	if err := r.db.Preload("User").First(&address, "id = ?", id).Error; err != nil {
		return model.Address{}, nil
	}

	if err := r.db.Model(&address).Updates(payload).Error; err != nil {
		return model.Address{}, err
	}

	return address, nil
}

// CreateAddress implements AddressRepository.
func (r *addressRepository) CreateAddress(address model.Address) (model.Address, error) {
	if address.UserID == "" {
		return model.Address{}, errors.New("user_id cannot be empty")
	}

	address.ID = utils.GenerateUuid()

	count, err := r.CountAddress(address.UserID)
	if err != nil {
		return model.Address{}, err
	}

	if count >= 5 {
		return model.Address{}, errors.New("user can only have 5 address")
	}

	if err := r.db.Create(&address).Error; err != nil {
		return model.Address{}, err
	}

	var savedAddress model.Address
	if err := r.db.Preload("User").First(&savedAddress, "id = ?", address.ID).Error; err != nil {
		return model.Address{}, err
	}

	return savedAddress, nil
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{
		db: db,
	}
}
