package repository

import (
	"github.com/hideUW/nuxt-go-chat-app/server/domain/model"
)

// SessionRepository is repository of session.
type SessionRepository interface {
	GetSessionByID(m SQLManager, id string) (*model.Session, error)
	InsertSession(m SQLManager, user *model.Session) error
	DeleteSession(m SQLManager, id string) error
}
