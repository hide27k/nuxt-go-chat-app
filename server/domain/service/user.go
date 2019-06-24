package service

import (
	"context"

	"github.com/hideUW/nuxt-go-chat-app/server/domain/repository"
	"github.com/pkg/errors"
)

// UserService is interface of domain service of user.
type UserService interface {
	IsAlreadyExistID(ctx context.Context, id uint32) (bool, error)
	IsAlreadyExistName(ctx context.Context, name string) (bool, error)
}

// UserRepoFactory is factory of UserRepository.
type UserRepoFactory func(ctx context.Context) repository.UserRepository

type userService struct {
	m    repository.DBManager
	repo repository.UserRepository
}

// NewUserService returns UserService.
func NewUserService(m repository.DBManager, repo repository.UserRepository) UserService {
	return &userService{
		m:    m,
		repo: repo,
	}
}

func (s *userService) IsAlreadyExistID(ctx context.Context, id uint32) (bool, error) {
	searched, err := s.repo.GetUserByID(s.m, id)
	if err != nil {
		return false, errors.Wrap(err, "failed to get user by id")
	}
	return searched != nil, nil
}

func (s *userService) IsAlreadyExistName(ctx context.Context, name string) (bool, error) {
	searched, err := s.repo.GetUserByName(s.m, name)
	if err != nil {
		return false, errors.Wrap(err, "failed to get user by name")
	}

	return searched != nil, nil
}
