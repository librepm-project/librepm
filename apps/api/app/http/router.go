package http

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	ProjectController      ProjectControllerInterface
	IssueController        IssueControllerInterface
	FilterController       FilterControllerInterface
	BoardController        BoardControllerInterface
	DashboardController    DashboardControllerInterface
	UserCurrentController  UserCurrentControllerInterface
	UserSessionController  UserSessionControllerInterface
	UserRegisterController UserRegisterControllerInterface
	UserController         UserControllerInterface
	TrackerController      TrackerControllerInterface
	StatusController       StatusControllerInterface
	TransitionController   TransitionControllerInterface
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

	router.Get("/dashboard", r.DashboardController.Index)
	router.Post("/dashboard", r.DashboardController.Create)
	router.Get("/dashboard/{dashboard_id}", r.DashboardController.Show)
	router.Put("/dashboard/{dashboard_id}", r.DashboardController.Update)
	router.Delete("/dashboard/{dashboard_id}", r.DashboardController.Destroy)

	router.Get("/user/current", r.UserCurrentController.Show)
	router.Put("/user/current", r.UserCurrentController.Update)
	router.Delete("/user/current", r.UserCurrentController.Destroy)
	router.Post("/user/register", r.UserRegisterController.Create)
	router.Post("/user/session", r.UserSessionController.Create)

	router.Get("/user", r.UserController.Index)
	router.Post("/user", r.UserController.Create)
	router.Get("/user/{user_id}", r.UserController.Show)
	router.Put("/user/{user_id}", r.UserController.Update)
	router.Delete("/user/{user_id}", r.UserController.Destroy)

	router.Get("/tracker", r.TrackerController.Index)
	router.Post("/tracker", r.TrackerController.Create)
	router.Get("/tracker/{tracker_id}", r.TrackerController.Show)
	router.Put("/tracker/{tracker_id}", r.TrackerController.Update)
	router.Delete("/tracker/{tracker_id}", r.TrackerController.Destroy)

	router.Get("/status", r.StatusController.Index)
	router.Post("/status", r.StatusController.Create)
	router.Get("/status/{status_id}", r.StatusController.Show)
	router.Put("/status/{status_id}", r.StatusController.Update)
	router.Delete("/status/{status_id}", r.StatusController.Destroy)

	router.Get("/transition", r.TransitionController.Index)
	router.Post("/transition", r.TransitionController.Create)
	router.Get("/transition/{transition_id}", r.TransitionController.Show)
	router.Put("/transition/{transition_id}", r.TransitionController.Update)
	router.Delete("/transition/{transition_id}", r.TransitionController.Destroy)

	return router
}
