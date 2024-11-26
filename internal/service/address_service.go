package service

import (
	"library_app/internal/repository"
	"library_app/model"
	"library_app/model/dto"
)

type AddressService interface {
	CreateAddress(address model.Address) (model.Address, error)
	UpdateAddress(id string, payload dto.AddressDto) (dto.AddressDto, error)
	GetAddress(id string) (model.Address, error)
	DeleteAddrees(id string) error
}

type addressService struct {
	addressRepo repository.AddressRepository
}

// DeleteAddrees implements AddressService.
func (s *addressService) DeleteAddrees(id string) error {
	return s.addressRepo.DeleteAddress(id)
}

// GetAddress implements AddressService.
func (s *addressService) GetAddress(id string) (model.Address, error) {

	address, err := s.addressRepo.GetAddress(id)
	if err != nil {
		return model.Address{}, err
	}

	return address, nil
}

// UpdateAddress implements AddressService.
func (s *addressService) UpdateAddress(id string, payload dto.AddressDto) (dto.AddressDto, error) {

	address, err := s.addressRepo.UpdateAddress(id, payload)
	if err != nil {
		return dto.AddressDto{}, err
	}

	payload = dto.AddressDto{
		PhoneNumber: address.PhoneNumber,
		Street:      address.Street,
		PostalCode:  address.PostalCode,
		City:        address.City,
		Country:     address.Country,
		State:       address.State,
	}

	return payload, nil
}

// CreateAddress implements AddressService.
func (s *addressService) CreateAddress(address model.Address) (model.Address, error) {
	return s.addressRepo.CreateAddress(address)
}

func NewAddressRepository(addressRepo repository.AddressRepository) AddressService {
	return &addressService{
		addressRepo: addressRepo,
	}
}
