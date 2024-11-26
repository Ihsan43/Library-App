package dto

type AddressDto struct {
	PhoneNumber string `json:"phone_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	State       string `json:"state"`
	Country     string `json:"country"`
}
