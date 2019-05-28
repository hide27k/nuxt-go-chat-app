package service

import (
	"context"

	"github.com/hideUW/nuxt_go_template/server/domain/repository"
)

// UserService is interface of domain service of user.
type UserService interface {
	IsAlreadyExistID(ctx context.Context, id uint32) (bool, error)
	IsAlreadyExistName(ctx context.Context, name string) (bool, error)
}

type userService struct {
	repo repository.UserRepository
	m    repository.SQLManager
}

// NewUserService returns UserService.
func NewUserService(repo repository.UserRepository, m repository.SQLManager) UserService {
	return &userService{
		repo: repo,
		m:    m,
	}
}

func (s *userService) IsAlreadyExistID(ctx context.Context, id uint32) (bool, error) {
	panic("implement me")
}

func (s *userService) IsAlreadyExistName(ctx context.Context, name string) (bool, error) {
	panic("implement me")
}
