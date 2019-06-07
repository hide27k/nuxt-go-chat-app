package application

import (
	"context"

	"github.com/hideUW/nuxt-go-chat-app/server/domain/model"
	"github.com/hideUW/nuxt-go-chat-app/server/domain/repository"
	"github.com/hideUW/nuxt-go-chat-app/server/domain/service"
)

// AuthenticationService is an interface of AuthenticationService.
type AuthenticationService interface {
	SignUp(ctx context.Context, name, password string) (*model.User, error)
}

type authenticationService struct {
	m        repository.SQLManager
	uFactory service.UserRepoFactory
	sFactory service.SessionRepoFactory
	txCloser CloseTransaction
}

// NewAuthenticationApplication returns AuthenticationService.
func NewAuthenticationApplication(m repository.SQLManager, uFactory service.UserRepoFactory, sFactory service.SessionRepoFactory, txCloser CloseTransaction) AuthenticationService {
	return &authenticationService{
		m:        m,
		uFactory: uFactory,
		sFactory: sFactory,
		txCloser: txCloser,
	}
}

func (a *authenticationService) SignUp(ctx context.Context, name, password string) (user *model.User, err error) {
	return nil, nil
}
