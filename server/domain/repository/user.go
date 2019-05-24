package repository

import "github.com/hideUW/nuxt_go_template/server/domain/model"

// UserRepository is repository of user.
type UserRepository interface {
	GetUserByID(m DBManager, id uint32) (*model.User, error)
	GetUserByName(m DBManager, name string) (*model.User, error)
	InsertUser(m DBManager, user *model.User) (uint32, error)
	UpdateUser(m DBManager, id uint32, user *model.User) error
	DeleteUser(m DBManager, id uint32) error
}
