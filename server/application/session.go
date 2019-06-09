package application

import (
	"context"

	"github.com/hideUW/nuxt-go-chat-app/server/domain/model"
	"github.com/hideUW/nuxt-go-chat-app/server/domain/repository"
	"github.com/hideUW/nuxt-go-chat-app/server/domain/service"
	"github.com/pkg/errors"
)

// SessionService is the interface of SessionService.
type SessionService interface {
}

// sessionService is the service of user.
type sessionService struct {
	m        repository.SQLManager
	sFactory service.SessionRepoFactory
	txCloser CloseTransaction
}

// NewSessionService generates and returns NewSessionService.
func NewSessionService(m repository.SQLManager, f service.SessionRepoFactory, txCloser CloseTransaction) SessionService {
	return &sessionService{
		m:        m,
		sFactory: f,
		txCloser: txCloser,
	}
}

func createSession(ctx context.Context, sessionID string, userID uint32, db repository.DBManager, repo repository.SessionRepository, sService service.SessionService) (*model.Session, error) {
	session := model.NewSession(userID)
	session.ID = sessionID

	yes := true
	var err error
	for yes {
		yes, err = sService.IsAlreadyExistID(ctx, session.ID)
		if err != nil {
			if _, ok := errors.Cause(err).(*model.NoSuchDataError); !ok {
				return nil, errors.Wrap(err, "failed to check whether already exists id or not")
			}
		}
	}

	if err := repo.InsertSession(db, session); err != nil {
		return nil, errors.Wrap(err, "failed to insert session")
	}
	return session, nil
}
