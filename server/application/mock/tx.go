package mock_application

import "github.com/hideUW/nuxt-go-chat-app/server/domain/repository"

// MockCloseTransaction executes after process of tx.
func MockCloseTransaction(tx repository.TxManager, err error) error {
	return nil
}
