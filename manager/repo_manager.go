package manager

import "library_app/internal/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	BookRepo() repository.BookRepository
}

type repoManager struct {
	infra InfraManager
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
