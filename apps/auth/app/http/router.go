package http

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	OAuth2Controller OAuth2ControllerInterface
}

func (r Router) Init() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/authorize", r.OAuth2Controller.Authorize)
	router.Get("/token", r.OAuth2Controller.Token)

	return router
}
