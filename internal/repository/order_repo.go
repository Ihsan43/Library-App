package repository

import (
	"errors"
	"fmt"
	"library_app/model"
	"library_app/utils"
	"time"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(payload model.Orders) (model.Orders, error)
	GetOrder(id string) (model.Orders, error)
}

type orderRepository struct {
	db *gorm.DB
}

// CreateOrder implements OrderRepository.
func (o *orderRepository) CreateOrder(payload model.Orders) (model.Orders, error) {
	var book model.Book
	if err := o.db.First(&book, "id = ?", payload.BookID).Error; err != nil {
		return model.Orders{}, fmt.Errorf("book not found: %v", err)
	}

	var expiredOrders []model.Orders
	if err := o.db.Where("expiry_time <= ?", time.Now()).Find(&expiredOrders).Error; err != nil {
		return model.Orders{}, fmt.Errorf("failed to fetch expired orders: %v", err)
	}

	for _, expiredOrder := range expiredOrders {
		if err := o.db.Where("order_id = ?", expiredOrder.ID).Delete(&model.Payment{}).Error; err != nil {
			return model.Orders{}, fmt.Errorf("failed to clean up payments for expired orders: %v", err)
		}
	}

	payload.ID = utils.GenerateUuid()
	payload.OrderDate = time.Now()
	payload.ExpiryTime = time.Now().Add(24 * time.Hour)

	if err := o.db.Where("expiry_time <= ?", time.Now()).Delete(&model.Orders{}).Error; err != nil {
		return model.Orders{}, fmt.Errorf("failed to clean up expired orders: %v", err)
	}

	payload.TotalPrice = book.Price * payload.Quantity

	if payload.Quantity > book.StockAvailable {
		return model.Orders{}, errors.New("insufficient stock for the order")
	}

	if err := o.db.Create(&payload).Error; err != nil {
		return model.Orders{}, fmt.Errorf("failed to create order: %v", err)
	}

	var createdOrder model.Orders
	if err := o.db.Preload("User").Preload("Book").Preload("ShippingAddress").First(&createdOrder, "id = ?", payload.ID).Error; err != nil {
		return model.Orders{}, fmt.Errorf("failed to fetch created order: %v", err)
	}

	return createdOrder, nil
}

// GetOrder implements OrderRepository.
func (o *orderRepository) GetOrder(id string) (model.Orders, error) {
	panic("unimplemented")
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}
