package manager

import "library_app/internal/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infra InfraManager
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
