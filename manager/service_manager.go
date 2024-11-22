package manager

import "library_app/internal/service"

type ServiceManager interface {
	UserService() service.UserService
	AuthService() service.AuthService
}

type serviceManager struct {
	repo RepoManager
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
