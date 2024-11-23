package http

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	ProjectController ProjectControllerInterface
}

func (r Router) Init() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/project", r.ProjectController.Index)
	router.Post("/project", r.ProjectController.Create)
	router.Get("/project/{project_id}", r.ProjectController.Show)
	router.Put("/project/{project_id}", r.ProjectController.Update)
	router.Delete("/project/{project_id}", r.ProjectController.Destroy)
	return router
}
