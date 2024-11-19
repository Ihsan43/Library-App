package manager

import "library_app/internal/repository"

type RepoManager interface {
	AccountRepo() repository.AccountRepository
}

type repoManager struct {
	infra InfraManager
}

// UserRepo implements RepoManager.
func (r *repoManager) AccountRepo() repository.AccountRepository {
	return repository.NewAccountRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
