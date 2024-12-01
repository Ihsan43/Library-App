package repository

import (
	"errors"
	"fmt"
	"library_app/model"
	"library_app/utils"

	"gorm.io/gorm"
)

type PaymentRepo interface {
	CreatePayment(payment model.Payment) (model.Payment, error)
	GetPaymentByOrderID(orderID string) (*model.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

// CreatePayment implements PaymentRepo.
func (r *paymentRepository) CreatePayment(payment model.Payment) (model.Payment, error) {
	// Preload data order terkait untuk validasi
	var order model.Orders
	if err := r.db.Preload("User").Preload("Book").First(&order, "id = ?", payment.OrderID).Error; err != nil {
		return model.Payment{}, fmt.Errorf("order not found: %v", err)
	}

	// Validasi jumlah pembayaran
	if payment.PayOrder < order.TotalPrice {
		return model.Payment{}, errors.New("payment amount is less than the total order price")
	}

	// Preload data buku terkait untuk validasi stok
	var book model.Book
	if err := r.db.First(&book, "id = ?", order.BookID).Error; err != nil {
		return model.Payment{}, fmt.Errorf("book not found: %v", err)
	}

	// Validasi stok buku
	if order.Quantity > book.StockAvailable {
		return model.Payment{}, errors.New("insufficient stock for the order")
	}

	// Set status pembayaran
	payment.ID = utils.GenerateUuid() // Generate ID untuk payment baru

	payment.Status = "Success"
	// Simpan payment ke database
	if err := r.db.Create(&payment).Error; err != nil {
		return model.Payment{}, fmt.Errorf("failed to create payment: %v", err)
	}

	// Update status order menjadi "Success"
	if err := r.db.Model(&order).Where("id = ?", payment.OrderID).Update("status", "Success").Error; err != nil {
		return model.Payment{}, fmt.Errorf("failed to update order status: %v", err)
	}

	// Kurangi stok buku
	newStock := book.StockAvailable - order.Quantity
	if err := r.db.Model(&book).Where("id = ?", order.BookID).Update("stock_available", newStock).Error; err != nil {
		return model.Payment{}, fmt.Errorf("failed to update book stock: %v", err)
	}

	return payment, nil
}

// GetPaymentByOrderID implements PaymentRepo.
func (p *paymentRepository) GetPaymentByOrderID(orderID string) (*model.Payment, error) {
	panic("unimplemented")
}

func NewPaymentRepository(db *gorm.DB) PaymentRepo {
	return &paymentRepository{
		db: db,
	}
}
