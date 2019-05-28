package application

import (
	"github.com/hideUW/nuxt_go_template/server/domain/model"
	"github.com/hideUW/nuxt_go_template/server/domain/repository"
	"github.com/pkg/errors"
)

// CloseTransaction work after finishing tx.
type CloseTransaction func(tx repository.TxManager, err error) error

func beginTxErrorMsg(err error) error {
	return errors.WithStack(&model.SQLError{
		BaseErr:                   err,
		InvalidReasonForDeveloper: model.FailedToBeginTx,
	})
}
