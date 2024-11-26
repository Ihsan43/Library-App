package manager

import "library_app/internal/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	BookRepo() repository.BookRepository
	AddressRepo() repository.AddressRepository
}

type repoManager struct {
	infra InfraManager
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
