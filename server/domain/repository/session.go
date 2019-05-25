package repository

import (
	"github.com/hideUW/nuxt_go_template/server/domain/model"
)

// SessionRepository is repository of session.
type SessionRepository interface {
	GetSessionByID(m DBManager, id string) (*model.Session, error)
	InsertSession(m DBManager, user *model.Session) error
	DeleteSession(m DBManager, id string) error
}
