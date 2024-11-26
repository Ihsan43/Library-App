package repository

import (
	"errors"
	"library_app/model"
	"library_app/model/dto"
	"library_app/utils"

	"gorm.io/gorm"
)

type AddressRepository interface {
	CreateAddress(address model.Address) (model.Address, error)
	UpdateAddress(id string, payload dto.AddressDto) (model.Address, error)
	DeleteAddress(id string) error
	GetAddress(id string) (model.Address, error)
}

type addressRepository struct {
	db *gorm.DB
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
func (r *addressRepository) UpdateAddress(id string, payload dto.AddressDto) (model.Address, error) {
	var address model.Address

	if err := r.db.First(&address, "id = ?", id).Error; err != nil {
		return model.Address{}, nil
	}

	if err := r.db.Model(&address).Updates(payload).Error; err != nil {
		return model.Address{}, err
	}

	return address, nil
}

// CreateAddress implements AddressRepository.
func (r *addressRepository) CreateAddress(address model.Address) (model.Address, error) {
	// Pastikan address memiliki UserID yang valid
	if address.UserID == "" {
		return model.Address{}, errors.New("user_id cannot be empty")
	}

	// Menambahkan ID baru untuk address
	address.ID = utils.GenerateUuid()

	// Menyimpan data address ke database
	if err := r.db.Create(&address).Error; err != nil {
		return model.Address{}, err
	}

	// Mengambil kembali address yang baru disimpan beserta User yang terkait
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
