package manager

import (
	"library_app/internal/service"
)

type ServiceManager interface {
	UserService() service.UserService
	AuthService() service.AuthService
	BookService() service.BookService
}

type serviceManager struct {
	repo RepoManager
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
