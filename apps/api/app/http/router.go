package http

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	ProjectController              ProjectControllerInterface
	IssueController                IssueControllerInterface
	IssueWorklogController         IssueWorklogControllerInterface
	RelatedIssueController         RelatedIssueControllerInterface
	FilterController               FilterControllerInterface
	BoardController                BoardControllerInterface
	DashboardController            DashboardControllerInterface
	DashboardWidgetController      DashboardWidgetControllerInterface
	UserCurrentController          UserCurrentControllerInterface
	UserSessionController          UserSessionControllerInterface
	UserRegisterController         UserRegisterControllerInterface
	UserController                 UserControllerInterface
	TrackerController              TrackerControllerInterface
	StatusController               StatusControllerInterface
	PriorityController             PriorityControllerInterface
	TransitionController           TransitionControllerInterface
	ProjectIssuePropertyController ProjectIssuePropertyController
	AttachmentController           AttachmentControllerInterface
	IssueAuditLogController        IssueAuditLogControllerInterface
	CommentController              CommentControllerInterface
	SettingController              SettingControllerInterface
	NotificationController         NotificationControllerInterface
	WebSocketController            WebSocketControllerInterface
}

func (r Router) Init() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/project", r.ProjectController.Index)
	router.Post("/project", r.ProjectController.Create)
	router.Get("/project/{project_id}", r.ProjectController.Show)
	router.Put("/project/{project_id}", r.ProjectController.Update)
	router.Delete("/project/{project_id}", r.ProjectController.Destroy)
	router.Get("/project/{project_id}/issue-property", r.ProjectIssuePropertyController.Index)
	router.Get("/project/{project_id}/tracker", r.ProjectController.IndexTrackers)
	router.Get("/project/{project_id}/status", r.ProjectController.IndexStatuses)

	router.Get("/issue", r.IssueController.Index)
	router.Post("/issue", r.IssueController.Create)
	router.Get("/issue/key/{key}", r.IssueController.ShowByKey)
	router.Get("/issue/{issue_id}", r.IssueController.Show)
	router.Put("/issue/{issue_id}", r.IssueController.Update)
	router.Delete("/issue/{issue_id}", r.IssueController.Destroy)
	router.Get("/issue/{issue_id}/related", r.RelatedIssueController.GetRelated)
	router.Post("/issue/{issue_id}/related", r.RelatedIssueController.Create)

	router.Get("/issue/{issue_id}/worklog", r.IssueWorklogController.Index)
	router.Post("/issue/{issue_id}/worklog", r.IssueWorklogController.Create)
	router.Put("/issue/{issue_id}/worklog/{worklog_id}", r.IssueWorklogController.Update)
	router.Delete("/issue/{issue_id}/worklog/{worklog_id}", r.IssueWorklogController.Destroy)

	router.Get("/issue/{issue_id}/attachment", r.AttachmentController.IndexByIssue)
	router.Post("/issue/{issue_id}/attachment", r.AttachmentController.CreateForIssue)
	router.Get("/project/{project_id}/attachment", r.AttachmentController.IndexByProject)
	router.Post("/project/{project_id}/attachment", r.AttachmentController.CreateForProject)
	router.Get("/attachment/{attachment_id}/download", r.AttachmentController.Download)
	router.Delete("/attachment/{attachment_id}", r.AttachmentController.Destroy)

	router.Get("/issue/{issue_id}/audit-log", r.IssueAuditLogController.Index)

	router.Get("/issue/{issue_id}/comment", r.CommentController.IndexByIssue)
	router.Post("/issue/{issue_id}/comment", r.CommentController.Create)
	router.Put("/comment/{comment_id}", r.CommentController.Update)
	router.Delete("/comment/{comment_id}", r.CommentController.Destroy)

	router.Delete("/related-issue/{related_issue_id}", r.RelatedIssueController.Delete)

	router.Get("/filter/condition-options", r.FilterController.Options)
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

	router.Get("/dashboard/{dashboard_id}/widget", r.DashboardWidgetController.Index)
	router.Post("/dashboard/{dashboard_id}/widget", r.DashboardWidgetController.Create)
	router.Put("/dashboard/{dashboard_id}/widget/reorder", r.DashboardWidgetController.Reorder)
	router.Patch("/dashboard/{dashboard_id}/widget/{widget_id}", r.DashboardWidgetController.Update)
	router.Delete("/dashboard/{dashboard_id}/widget/{widget_id}", r.DashboardWidgetController.Destroy)

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

	router.Get("/priority", r.PriorityController.Index)
	router.Post("/priority", r.PriorityController.Create)
	router.Get("/priority/{priority_id}", r.PriorityController.Show)
	router.Put("/priority/{priority_id}", r.PriorityController.Update)
	router.Delete("/priority/{priority_id}", r.PriorityController.Destroy)

	router.Get("/transition", r.TransitionController.Index)
	router.Post("/transition", r.TransitionController.Create)
	router.Get("/transition/{transition_id}", r.TransitionController.Show)
	router.Put("/transition/{transition_id}", r.TransitionController.Update)
	router.Delete("/transition/{transition_id}", r.TransitionController.Destroy)

	router.Get("/setting", r.SettingController.Index)
	router.Put("/setting/{key}", r.SettingController.Update)

	router.Get("/notification", r.NotificationController.Index)
	router.Patch("/notification/{notification_id}/read", r.NotificationController.MarkRead)

	router.Get("/ws", r.WebSocketController.Connect)

	return router
}
