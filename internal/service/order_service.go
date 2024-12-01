package service

import (
	"library_app/internal/repository"
	"library_app/model"
)

type OrderService interface {
	CreateOrder(payload model.Orders) (model.Orders, error)
}

type orderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{orderRepo: orderRepo}
}

func (o *orderService) CreateOrder(payload model.Orders) (model.Orders, error) {
	
	if payload.Status == "" {
		payload.Status = "Pending"
	}

	payload.TotalPrice = payload.Quantity * payload.Book.Price

	return o.orderRepo.CreateOrder(payload)
}
