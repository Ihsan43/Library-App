package service

import (
	"library_app/internal/repository"
	"library_app/model"
	"library_app/model/dto"
)

type AddressService interface {
	CreateAddress(payload dto.AddressRequestDto) (dto.AddressResponseDto, error)
	UpdateAddress(id string, payload dto.AddressRequestDto) (dto.AddressResponseDto, error)
	FindAddressByUserId(userId string) (model.Address, error)
	GetAddress(id string) (dto.AddressResponseDto, error)
	DeleteAddrees(id string) error
	FindAddresses(userId string) ([]dto.AddressResponseDto, error)
}

type addressService struct {
	addressRepo repository.AddressRepository
}

// FindAddresses implements AddressService.
func (s *addressService) FindAddresses(userId string) ([]dto.AddressResponseDto, error) {

	addresses, err := s.addressRepo.FindAddresses(userId)
	if err != nil {
		return nil, err
	}

	var addressDtos []dto.AddressResponseDto
	for _, address := range addresses {
		addressDtos = append(addressDtos, dto.NewAddressResponseDto(address))
	}

	return addressDtos, nil
}

// FindAddressByUserId implements AddressService.
func (s *addressService) FindAddressByUserId(userId string) (model.Address, error) {
	return s.addressRepo.FindAddressByUserId(userId)
}

// DeleteAddrees implements AddressService.
func (s *addressService) DeleteAddrees(id string) error {
	return s.addressRepo.DeleteAddress(id)
}

// GetAddress implements AddressService.
func (s *addressService) GetAddress(id string) (dto.AddressResponseDto, error) {

	address, err := s.addressRepo.GetAddress(id)
	if err != nil {
		return dto.AddressResponseDto{}, err
	}

	return dto.NewAddressResponseDto(address), nil
}

// UpdateAddress implements AddressService.
func (s *addressService) UpdateAddress(id string, payload dto.AddressRequestDto) (dto.AddressResponseDto, error) {

	address := model.Address{
		PhoneNumber: payload.PhoneNumber,
		Street:      payload.Street,
		City:        payload.City,
		PostalCode:  payload.PostalCode,
		State:       payload.State,
		Country:     payload.Country,
	}

	newAddress, err := s.addressRepo.UpdateAddress(id, address)
	if err != nil {
		return dto.AddressResponseDto{}, err
	}

	return dto.NewAddressResponseDto(newAddress), nil
}

// CreateAddress implements AddressService.
func (s *addressService) CreateAddress(payload dto.AddressRequestDto) (dto.AddressResponseDto, error) {

	address := model.Address{
		PhoneNumber: payload.PhoneNumber,
		Street:      payload.Street,
		City:        payload.City,
		UserID:      payload.UserId,
		PostalCode:  payload.PostalCode,
		State:       payload.State,
		Country:     payload.Country,
	}

	newAddress, err := s.addressRepo.CreateAddress(address)
	if err != nil {
		return dto.AddressResponseDto{}, err
	}

	return dto.NewAddressResponseDto(newAddress), nil
}

func NewAddressRepository(addressRepo repository.AddressRepository) AddressService {
	return &addressService{
		addressRepo: addressRepo,
	}
}
