package dto

import (
	"library_app/model"
	"time"
)

type AddressRequestDto struct {
	PhoneNumber string `json:"phone_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	UserId      string `json:"user_id"`
	PostalCode  string `json:"postal_code"`
	State       string `json:"state"`
	Country     string `json:"country"`
}

type AddressResponseDto struct {
	ID          string           `json:"id"`
	UserID      string           `json:"user_id"`
	User        *UserResponseDto `json:"user"`
	PhoneNumber string           `json:"phone_number"`
	Street      string           `json:"street"`
	City        string           `json:"city"`
	PostalCode  string           `json:"postal_code"`
	State       string           `json:"state"`
	Country     string           `json:"country"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

func NewAddressResponseDto(address model.Address) AddressResponseDto {
	return AddressResponseDto{
		ID:          address.ID,
		UserID:      address.UserID,
		User: &UserResponseDto{
			ID:          address.User.ID,
			Name:        address.User.Name,
			Email:       address.User.Email,
			PhoneNumber: address.User.PhoneNumber,
			Username:    address.User.Username,
			AvatarURL:   address.User.AvatarURL,
			CreatedAt:   address.User.CreatedAt,
			UpdatedAt:   address.User.UpdatedAt,
		},
		PhoneNumber: address.PhoneNumber,
		Street:      address.Street,
		City:        address.City,
		PostalCode:  address.PostalCode,
		State:       address.State,
		Country:     address.Country,
		CreatedAt:   address.CreatedAt,
		UpdatedAt:   address.UpdatedAt,
	}
}


