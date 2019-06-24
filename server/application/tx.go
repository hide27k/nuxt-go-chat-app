package application

import (
	"github.com/hideUW/nuxt-go-chat-app/server/domain/model"
	"github.com/hideUW/nuxt-go-chat-app/server/domain/repository"
	"github.com/pkg/errors"
)

// CloseTransaction executes after process of tx.
type CloseTransaction func(tx repository.TxManager, err error) error

// beginTxErrorMsg generates and returns tx begin error message.
func beginTxErrorMsg(err error) error {
	return errors.WithStack(&model.SQLError{
		BaseErr:                   err,
		InvalidReasonForDeveloper: model.FailedToBeginTx,
	})
}
