package http

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	ProjectController ProjectControllerInterface
}

func (r Router) Init() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/product", r.ProjectController.Index)
	router.Post("/product", r.ProjectController.Create)
	router.Get("/product/{product_id}", r.ProjectController.Show)
	router.Put("/product/{product_id}", r.ProjectController.Update)
	return router
}
