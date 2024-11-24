package http

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	ProjectController ProjectControllerInterface
	IssueController   IssueControllerInterface
	FilterController  FilterControllerInterface
	BoardController   BoardControllerInterface
}

func (r Router) Init() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/project", r.ProjectController.Index)
	router.Post("/project", r.ProjectController.Create)
	router.Get("/project/{project_id}", r.ProjectController.Show)
	router.Put("/project/{project_id}", r.ProjectController.Update)
	router.Delete("/project/{project_id}", r.ProjectController.Destroy)

	router.Get("/issue", r.IssueController.Index)
	router.Post("/issue", r.IssueController.Create)
	router.Get("/issue/{issue_id}", r.IssueController.Show)
	router.Put("/issue/{issue_id}", r.IssueController.Update)
	router.Delete("/issue/{issue_id}", r.IssueController.Destroy)

	router.Get("/filter", r.FilterController.Index)
	router.Post("/filter", r.FilterController.Create)
	router.Get("/filter/{filter_id}", r.FilterController.Show)
	router.Put("/filter/{filter_id}", r.FilterController.Update)
	router.Delete("/filter/{filter_id}", r.FilterController.Destroy)

	router.Get("/board", r.BoardController.Index)
	router.Post("/board", r.BoardController.Create)
	router.Get("/board/{board_id}", r.BoardController.Show)
	router.Put("/board/{board_id}", r.BoardController.Update)
	router.Delete("/board/{board_id}", r.BoardController.Destroy)

	return router
}
