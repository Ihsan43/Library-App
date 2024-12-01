package model

import "time"

type Orders struct {
	ID                string    `gorm:"primary_key;type:uuid;" json:"id"`
	UserID            string    `gorm:"type:uuid;not null" json:"user_id"`
	User              *User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	BookID            string    `gorm:"type:uuid;not null" json:"book_id"`
	Book              *Book      `gorm:"foreignKey:BookID;references:ID" json:"book"`
	Quantity          int       `gorm:"size:255;not null" json:"quantity"`
	TotalPrice        int       `gorm:"size:255;not null" json:"total_price"`
	Status            string    `gorm:"size:255;not null" json:"status"`
	ShippingAddressID string    `gorm:"type:uuid;not null" json:"shipping_address_id"`
	ShippingAddress   *Address   `gorm:"foreignKey:ShippingAddressID;references:ID" json:"shipping_address"`
	OrderDate         time.Time `gorm:"default:current_timestamp" json:"order_date"`
	ExpiryTime        time.Time `json:"expiry_time"` // Waktu Kedaluwarsa
	CreatedAt         time.Time `gorm:"default:current_timestamp" json:"created_at"`
}
