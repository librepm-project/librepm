package http

import (
	"apps/auth/app/domain"
	"net/http"
)

type OAuth2ControllerInterface interface {
	Authorize(w http.ResponseWriter, r *http.Request)
	Token(w http.ResponseWriter, r *http.Request)
}

type OAuth2Controller struct {
	OAuth2Service domain.OAuth2ServiceInterface
}

func (c OAuth2Controller) Authorize(w http.ResponseWriter, r *http.Request) {
	err := c.OAuth2Service.Get().HandleAuthorizeRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (c OAuth2Controller) Token(w http.ResponseWriter, r *http.Request) {
	c.OAuth2Service.Get().HandleTokenRequest(w, r)
}
