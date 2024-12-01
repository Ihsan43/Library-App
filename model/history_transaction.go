package model

import "time"

type TransactionHistory struct {
	ID              string    `gorm:"primary_key;type:uuid;" json:"id"`
	PaymentID       string    `gorm:"type:uuid;not null" json:"payment_id"`
	Payment         Payment   `gorm:"foreignKey:PaymentID;references:ID" json:"payment"`
	OrderID         string    `gorm:"type:uuid;not null" json:"order_id"`
	Order           Orders    `gorm:"foreignKey:OrderID;references:ID" json:"order"`
	Amount          int       `gorm:"not null" json:"amount"`                    // Jumlah transaksi
	Status          string    `gorm:"size:100;not null" json:"status"`           // e.g., "Success", "Failed"
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`          // Waktu transaksi dibuat
}
