package model

import "time"

type Payment struct {
	ID        string    `gorm:"primary_key;type:uuid" json:"id"`
	OrderID   string    `gorm:"not null" json:"order_id"`
	Order     *Orders    `gorm:"foreignKey:OrderID;references:ID" json:"-"`
	PayOrder  int       `gorm:"not null" json:"pay_order"`
	Status    string    `gorm:"size:50;not null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:null" json:"updatedAt"`
}
