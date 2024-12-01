package manager

import (
	"library_app/internal/service"
)

type ServiceManager interface {
	UserService() service.UserService
	AuthService() service.AuthService
	BookService() service.BookService
	AddressService() service.AddressService
	OrderService() service.OrderService
	PaymentService() service.PaymentService
	TransactionService() service.TransactionService
}

type serviceManager struct {
	repo RepoManager
}

// TransactionService implements ServiceManager.
func (s *serviceManager) TransactionService() service.TransactionService {
	return service.NewTransactionService(s.repo.TransactionRepo())
}

// PaymentService implements ServiceManager.
func (s *serviceManager) PaymentService() service.PaymentService {
	return service.NewPaymentService(s.repo.PaymentRepo())
}

// OrderService implements ServiceManager.
func (s *serviceManager) OrderService() service.OrderService {
	return service.NewOrderService(s.repo.OrderRepo())
}

// AddressService implements ServiceManager.
func (s *serviceManager) AddressService() service.AddressService {
	return service.NewAddressRepository(s.repo.AddressRepo())
}

// BookService implements ServiceManager.
func (s *serviceManager) BookService() service.BookService {
	return service.NewBookService(s.repo.BookRepo())
}

// AccountService implements ServiceManager.
func (s *serviceManager) UserService() service.UserService {
	return service.NewAccountService(s.repo.UserRepo())
}

// AuthService implements ServiceManager.
func (s *serviceManager) AuthService() service.AuthService {
	return service.NewAuthService(s.UserService())
}

func NewServiceManager(repo RepoManager) ServiceManager {
	return &serviceManager{repo: repo}
}
