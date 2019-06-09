package model

import (
	"time"

	"github.com/hideUW/nuxt-go-chat-app/server/util"
)

// User is User model
type User struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	SessionID string    `json:"sessionId"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewUser creates and returns User model.
func NewUser(name, password string) (*User, error) {
	hashed, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		Password: hashed,
	}, nil
}
