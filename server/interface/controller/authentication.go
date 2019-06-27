package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/hideUW/nuxt-go-chat-app/server/application"
	"github.com/hideUW/nuxt-go-chat-app/server/domain/model"
	"github.com/hideUW/nuxt-go-chat-app/server/infra/router"
	"github.com/pkg/errors"
)

// AuthenticationController is the interface of AuthenticationController.
type AuthenticationController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
}

type authenticationController struct {
	rm   router.RequestManager
	aApp application.AuthenticationService
}

// NewAuthenticationController generates and returns AuthenticationController.
func NewAuthenticationController(rm router.RequestManager, uAPP application.AuthenticationService) AuthenticationController {
	return &authenticationController{
		rm:   rm,
		aApp: uAPP,
	}
}

func (c *authenticationController) SignUp(w http.ResponseWriter, r *http.Request) {
	b, err := GetValueFromPayLoad(r)
	if err != nil {
		ResponseAndLogError(w, err)
		return
	}

	user, err := ParseUserFromPayLoad(b)
	if err != nil {
		ResponseAndLogError(w, err)
		return
	}

	user, err = model.NewUser(user.Name, user.Password)
	if err != nil {
		ResponseAndLogError(w, err)
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	ctx := r.Context()
	user, err = c.aApp.SignUp(ctx, user)
	if err != nil {
		ResponseAndLogError(w, err)
		return
	}

	cookie := c.newCookieWithSessionID(user.SessionID, 86400)
	uDTO := TranslateFromUserToUserDTO(user)

	if err := ResponseWithCookie(w, http.StatusOK, cookie, uDTO); err != nil {
		ResponseAndLogError(w, err)
		return
	}
}

func ParseUserFromPayload(b []byte) (*model.User, error) {
	u := &model.User{}
	if err := json.Unmarshal(b, u); err != nil {
		if err := json.Unmarshal(b, u); err != nil {
			err = &model.InvalidDataError{
				BaseErr:               err,
				DataNameForDeveloper:  "request body",
				DataValueForDeveloper: string(b),
			}
			return nil, errors.WithStack(err)
		}
	}
	return u, nil
}

// newCookieWithSessionID generates and returns cookie with session id.
func (c *authenticationController) newCookieWithSessionID(sessionID string, maxAge int) *http.Cookie {
	return &http.Cookie{
		Name:   model.SessionIDAtCookie,
		Value:  sessionID,
		MaxAge: maxAge,
	}
}
