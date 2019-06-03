package repository

import (
	"github.com/hideUW/nuxt-go-chat-app/server/domain/model"
)

// SessionRepository is repository of session.
type SessionRepository interface {
	GetSessionByID(m DBManager, id string) (*model.Session, error)
	InsertSession(m DBManager, user *model.Session) error
	DeleteSession(m DBManager, id string) error
}
