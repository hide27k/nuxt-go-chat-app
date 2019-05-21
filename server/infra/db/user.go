package db

import (
	"context"

	"github.com/hideUW/nuxt_go_template/server/domain/model"
	. "github.com/hideUW/nuxt_go_template/server/domain/repository"
)

// userRepository is the repository of the user.
type userRepository struct {
	ctx context.Context
}

// NewUserRepository generates userRepository.
func NewUserRepository(ctx context.Context) UserRepository {
	return &userRepository{
		ctx: ctx,
	}
}

func (repo *userRepository) ErrorMsg(method RepositoryMethod, err error) error {
	return nil
}

func (repo *userRepository) GetUserByID(m DBManager, id uint32) (*model.User, error) {
	return nil, nil
}
func (repo *userRepository) GetUserByName(m DBManager, name string) (*model.User, error) {
	return nil, nil
}
func (repo *userRepository) InsertUser(m DBManager, user *model.User) (uint32, error) {
	return 0, nil
}
func (repo *userRepository) UpdateUser(m DBManager, id uint32, user *model.User) error {
	return nil
}
func (repo *userRepository) DeleteUser(m DBManager, id uint32) error {
	return nil
}
