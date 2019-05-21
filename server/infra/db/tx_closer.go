package db

import (
	"github.com/hideUW/nuxt_go_template/server/domain/repository"
	"github.com/pkg/errors"
)

// CloseTransaction executes post process of tx
func CloseTransaction(tx repository.TxManager, err error) error {
	if p := recover(); p != nil {
		err = tx.Rollback()
		panic(p)
	} else if err != nil {
		err = tx.Rollback()
	} else {
		err = tx.Commit()
	}
	return errors.WithStack(err)
}
