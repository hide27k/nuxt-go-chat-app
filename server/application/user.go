package application

import (
	"github.com/hideUW/nuxt-go-chat-app/server/domain/repository"
	"github.com/hideUW/nuxt-go-chat-app/server/domain/service"
)

// UserService is an interface.
type UserService interface {
}

type userService struct {
	m        repository.DBManager
	uFactory service.UserRepoFactory
	txCloser CloseTransaction
}

// NewUserService creates an interface called UserService and returns it.
func NewUserService(m repository.DBManager, f service.UserRepoFactory, txCloser CloseTransaction) UserService {
	return &userService{
		m:        m,
		uFactory: f,
		txCloser: txCloser,
	}
}
