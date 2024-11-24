package dto

type UserDto struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	AvatarURL   string `json:"avatar_url"`
}
