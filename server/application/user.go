package application

import (
	"context"
	"time"

	"github.com/hideUW/nuxt-go-chat-app/server/domain/model"
	"github.com/hideUW/nuxt-go-chat-app/server/domain/repository"
	"github.com/hideUW/nuxt-go-chat-app/server/domain/service"
	"github.com/pkg/errors"
)

// UserService is an interface.
type UserService interface {
}

type userService struct {
	m        repository.SQLManager
	uFactory service.UserRepoFactory
	txCloser CloseTransaction
}

// NewUserService creates an interface called UserService and returns it.
func NewUserService(m repository.SQLManager, f service.UserRepoFactory, txCloser CloseTransaction) UserService {
	return &userService{
		m:        m,
		uFactory: f,
		txCloser: txCloser,
	}
}

func createUser(ctx context.Context, param *model.User, db repository.DBManager, repo repository.UserRepository, uService service.UserService) (*model.User, error) {
	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()

	yes, err := uService.IsAlreadyExistName(ctx, param.Name)
	if yes {
		err = &model.AlreadyExistError{
			PropertyNameForDeveloper:    model.NamePropertyForDeveloper,
			PropertyNameForUser:         model.NamePropertyForUser,
			PropertyValue:               param.Name,
			DomainModelNameForDeveloper: model.DomainModelNameUserForDeveloper,
			DomainModelNameForUser:      model.DomainModelNameUserForUser,
		}

		return nil, errors.Wrap(err, "failed to check if the name already exists or not")
	}

	if _, ok := errors.Cause(err).(*model.NoSuchDataError); !ok {
		return nil, errors.Wrap(err, "failed to check whether already exists name or not")
	}

	id, err := repo.InsertUser(db, param)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert user")
	}
	param.ID = id

	return param, nil
}
