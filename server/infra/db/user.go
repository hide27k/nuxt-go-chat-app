package db

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/hideUW/nuxt_go_template/server/domain/model"
	. "github.com/hideUW/nuxt_go_template/server/domain/repository"
	"github.com/pkg/errors"
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

func (repo *userRepository) ErrorMsg(method model.RepositoryMethod, err error) error {
	return &model.RepositoryError{
		BaseErr:                     err,
		RepositoryMethod:            method,
		DomainModelNameForDeveloper: model.DomainModelNameUserForDeveloper,
		DomainModelNameForUser:      model.DomainModelNameUserForUser,
	}
}

func (repo *userRepository) GetUserByID(m DBManager, id uint32) (*model.User, error) {
	query := "SELECT id, name, session_id, password, created_at, updated_at FROM users WHERE id=?"

	list, err := repo.list(m, model.RepositoryMethodREAD, query, id)

	if len(list) == 0 {
		err = &model.NoSuchDataError{
			BaseErr:                     err,
			PropertyNameForDeveloper:    model.IDPropertyForDeveloper,
			PropertyNameForUser:         model.IDPropertyForUser,
			PropertyValue:               id,
			DomainModelNameForDeveloper: model.DomainModelNameUserForDeveloper,
			DomainModelNameForUser:      model.DomainModelNameUserForUser,
		}
		return nil, err
	}

	if err != nil {
		return nil, errors.WithStack(repo.ErrorMsg(model.RepositoryMethodREAD, errors.WithStack(err)))
	}

	return list[0], nil
}

func (repo *userRepository) GetUserByName(m DBManager, name string) (*model.User, error) {
	query := "SELECT id, name, session_id, password, created_at, updated_at FROM users WHERE name=?"
	list, err := repo.list(m, model.RepositoryMethodREAD, query, name)

	if len(list) == 0 {
		err = &model.NoSuchDataError{
			BaseErr:                     err,
			PropertyNameForDeveloper:    model.NamePropertyForDeveloper,
			PropertyNameForUser:         model.NamePropertyForUser,
			PropertyValue:               name,
			DomainModelNameForDeveloper: model.DomainModelNameUserForDeveloper,
			DomainModelNameForUser:      model.DomainModelNameUserForUser,
		}
		return nil, err
	}

	if err != nil {
		return nil, errors.WithStack(repo.ErrorMsg(model.RepositoryMethodREAD, err))
	}

	return list[0], nil

}

func (repo *userRepository) list(m DBManager, method model.RepositoryMethod, query string, args ...interface{}) (users []*model.User, err error) {
	stmt, error := m.PrepareContext(repo.ctx, query)
	if err != nil {
		return nil, repo.ErrorMsg(method, err)
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Error(error.Error())
		}
	}()

	rows, err := stmt.QueryContext(repo.ctx, args...)
	if err != nil {
		return nil, errors.WithStack(repo.ErrorMsg(method, err))
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}()

	list := make([]*model.User, 0)
	for rows.Next() {
		user := &model.User{}

		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.SessionID,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, errors.WithStack(repo.ErrorMsg(method, err))
		}

		list = append(list, user)
	}

	return list, nil
}

func (repo *userRepository) InsertUser(m DBManager, user *model.User) (uint32, error) {
	query := "INSERT INFO users (name, session_id, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := m.PrepareContext(repo.ctx, query)
	if err != nil {
		return model.InvalidID, errors.WithStack(repo.ErrorMsg(model.RepositoryMethodINSERT, err))
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}()

	result, err := stmt.ExecContext(repo.ctx, user.Name, user.SessionID, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return model.InvalidID, errors.WithStack(repo.ErrorMsg(model.RepositoryMethodINSERT, err))
	}

	affect, err := result.RowsAffected()
	if affect != 1 {
		err = fmt.Errorf("total affected: %d ", affect)
		return model.InvalidID, errors.WithStack(repo.ErrorMsg(model.RepositoryMethodINSERT, err))
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.InvalidID, errors.WithStack(repo.ErrorMsg(model.RepositoryMethodINSERT, err))
	}

	return uint32(id), nil
}
func (repo *userRepository) UpdateUser(m DBManager, id uint32, user *model.User) error {
	return nil
}
func (repo *userRepository) DeleteUser(m DBManager, id uint32) error {
	return nil
}
