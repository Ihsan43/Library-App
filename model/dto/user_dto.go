package dto

import (
	"library_app/model"
	"time"
)

type UserRequestDto struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"`
	Role        string `json:"role"`
	AvatarURL   string `json:"avatar_url"`
}

type UserResponseDto struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Username    string    `json:"username"`
	Role        string    `json:"role"`
	IsVerified  bool      `json:"is_verified"`
	Status      string    `json:"status"`
	AvatarURL   string    `json:"avatar_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewUserResponseDto(user model.User) UserResponseDto {
	return UserResponseDto{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Username:    user.Username,
		Role:        user.Role,
		IsVerified:  user.IsVerified,
		Status:      user.Status,
		AvatarURL:   user.AvatarURL,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}
