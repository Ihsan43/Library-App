package service

import (
	"library_app/internal/repository"
	"library_app/model"
)

type PaymentService interface {
	CreatePayment(payment model.Payment) (model.Payment, error)
}

type paymentService struct {
	repo repository.PaymentRepo
}

// CreatePayment implements PaymentService.
func (p *paymentService) CreatePayment(payment model.Payment) (model.Payment, error) {
	return p.repo.CreatePayment(payment)
}

func NewPaymentService(repo repository.PaymentRepo) PaymentService {
	return &paymentService{
		repo: repo,
	}
}
