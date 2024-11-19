package manager

import "library_app/internal/service"

type ServiceManager interface {
	AccountService() service.AccountService
	AuthService() service.AuthService
}

type serviceManager struct {
	repo RepoManager
}

// AccountService implements ServiceManager.
func (s *serviceManager) AccountService() service.AccountService {
	return service.NewAccountService(s.repo.AccountRepo())
}

// AuthService implements ServiceManager.
func (s *serviceManager) AuthService() service.AuthService {
	return service.NewAuthService(s.AccountService())
}

func NewServiceManager(repo RepoManager) ServiceManager {
	return &serviceManager{repo: repo}
}
