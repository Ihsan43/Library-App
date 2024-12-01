package manager

import "library_app/internal/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	BookRepo() repository.BookRepository
	AddressRepo() repository.AddressRepository
	OrderRepo() repository.OrderRepository
	PaymentRepo() repository.PaymentRepo
	TransactionRepo() repository.TransactionRepository
}

type repoManager struct {
	infra InfraManager
}

// TransactionRepo implements RepoManager.
func (r *repoManager) TransactionRepo() repository.TransactionRepository {
	return repository.NewTransactionRepository(r.infra.Conn())
}

// PaymentRepo implements RepoManager.
func (r *repoManager) PaymentRepo() repository.PaymentRepo {
	return repository.NewPaymentRepository(r.infra.Conn())
}

// OrderRepo implements RepoManager.
func (r *repoManager) OrderRepo() repository.OrderRepository {
	return repository.NewOrderRepository(r.infra.Conn())
}

// AddressRepo implements RepoManager.
func (r *repoManager) AddressRepo() repository.AddressRepository {
	return repository.NewAddressRepository(r.infra.Conn())
}

// BookRepo implements RepoManager.
func (r *repoManager) BookRepo() repository.BookRepository {
	return repository.NewBookRepository(*r.infra.Conn())
}

// UserRepo implements RepoManager.
func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
