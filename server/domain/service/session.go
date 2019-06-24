package service

import (
	"context"

	"github.com/hideUW/nuxt-go-chat-app/server/domain/model"
	"github.com/hideUW/nuxt-go-chat-app/server/domain/repository"
	"github.com/hideUW/nuxt-go-chat-app/server/util"
	"github.com/pkg/errors"
)

// SessionService is inferface of domain service of session.
type SessionService interface {
	IsAlreadyExistID(ctx context.Context, id string) (bool, error)
	SessionID() string
}

// SessionRepoFactory creates of SessionRepository.
type SessionRepoFactory func(ctx context.Context) repository.SessionRepository

type sessionService struct {
	m    repository.DBManager
	repo repository.SessionRepository
}

// NewSessionService returns SessionService which is interface.
func NewSessionService(m repository.DBManager, repo repository.SessionRepository) SessionService {
	return &sessionService{
		m:    m,
		repo: repo,
	}
}

func (s *sessionService) SessionID() string {
	return util.UUID()
}

func (s sessionService) IsAlreadyExistID(ctx context.Context, id string) (bool, error) {
	var searched *model.Session
	var err error

	if searched, err = s.repo.GetSessionByID(s.m, id); err != nil {
		return false, errors.Wrap(err, "failed to get session by id")
	}

	return searched != nil, nil
}
