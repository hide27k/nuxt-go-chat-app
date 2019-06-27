package service

import (
	"context"
	"time"

	"github.com/hideUW/nuxt-go-chat-app/server/domain/model"
	"github.com/hideUW/nuxt-go-chat-app/server/domain/repository"
	"github.com/hideUW/nuxt-go-chat-app/server/util"
	"github.com/pkg/errors"
)

// UserService is interface of domain service of user.
type UserService interface {
	NewUser(name, password string) (*model.User, error)
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

// NewUser generates and reruns User.
func (s *userService) NewUser(name, password string) (*model.User, error) {
	hashed, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Name:      name,
		Password:  hashed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
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
