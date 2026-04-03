package http

import (
	"apps/api/app/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	OnboardController              OnboardControllerInterface
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
	ReleaseController              ReleaseControllerInterface
	ProjectReleaseController       ProjectReleaseControllerInterface
	ProjectReleaseIssueController  ProjectReleaseIssueControllerInterface
	PermissionService              domain.PermissionServiceInterface
	UserPermissionController       UserPermissionControllerInterface
	UserRoleController             UserRoleControllerInterface
}

func (r Router) Init() *chi.Mux {
	router := chi.NewRouter()

	router.With(RequirePermission(r.PermissionService, domain.ProjectRead)).Get("/project", r.ProjectController.Index)
	router.With(RequirePermission(r.PermissionService, domain.ProjectCreate)).Post("/project", r.ProjectController.Create)
	router.With(RequirePermission(r.PermissionService, domain.ProjectRead)).Get("/project/{project_id}", r.ProjectController.Show)
	router.With(RequirePermission(r.PermissionService, domain.ProjectUpdate)).Put("/project/{project_id}", r.ProjectController.Update)
	router.With(RequirePermission(r.PermissionService, domain.ProjectDelete)).Delete("/project/{project_id}", r.ProjectController.Destroy)
	router.With(RequirePermission(r.PermissionService, domain.ProjectRead)).Get("/project/{project_id}/issue-property", r.ProjectIssuePropertyController.Index)
	router.With(RequirePermission(r.PermissionService, domain.ProjectRead)).Get("/project/{project_id}/tracker", r.ProjectController.IndexTrackers)
	router.With(RequirePermission(r.PermissionService, domain.ProjectRead)).Get("/project/{project_id}/status", r.ProjectController.IndexStatuses)

	router.With(RequirePermission(r.PermissionService, domain.IssueRead)).Get("/issue", r.IssueController.Index)
	router.With(RequirePermission(r.PermissionService, domain.IssueCreate)).Post("/issue", r.IssueController.Create)
	router.With(RequirePermission(r.PermissionService, domain.IssueRead)).Get("/issue/key/{key}", r.IssueController.ShowByKey)
	router.With(RequirePermission(r.PermissionService, domain.IssueRead)).Get("/issue/{issue_id}", r.IssueController.Show)
	router.With(RequirePermission(r.PermissionService, domain.IssueUpdate)).Put("/issue/{issue_id}", r.IssueController.Update)
	router.With(RequirePermission(r.PermissionService, domain.IssueDelete)).Delete("/issue/{issue_id}", r.IssueController.Destroy)
	router.With(RequirePermission(r.PermissionService, domain.IssueRead)).Get("/issue/{issue_id}/related", r.RelatedIssueController.GetRelated)
	router.With(RequirePermission(r.PermissionService, domain.IssueUpdate)).Post("/issue/{issue_id}/related", r.RelatedIssueController.Create)

	router.With(RequirePermission(r.PermissionService, domain.WorklogRead)).Get("/issue/{issue_id}/worklog", r.IssueWorklogController.Index)
	router.With(RequirePermission(r.PermissionService, domain.WorklogCreate)).Post("/issue/{issue_id}/worklog", r.IssueWorklogController.Create)
	router.With(RequirePermission(r.PermissionService, domain.WorklogUpdate)).Put("/issue/{issue_id}/worklog/{worklog_id}", r.IssueWorklogController.Update)
	router.With(RequirePermission(r.PermissionService, domain.WorklogDelete)).Delete("/issue/{issue_id}/worklog/{worklog_id}", r.IssueWorklogController.Destroy)

	router.With(RequirePermission(r.PermissionService, domain.AttachmentRead)).Get("/issue/{issue_id}/attachment", r.AttachmentController.IndexByIssue)
	router.With(RequirePermission(r.PermissionService, domain.AttachmentCreate)).Post("/issue/{issue_id}/attachment", r.AttachmentController.CreateForIssue)
	router.With(RequirePermission(r.PermissionService, domain.AttachmentRead)).Get("/project/{project_id}/attachment", r.AttachmentController.IndexByProject)
	router.With(RequirePermission(r.PermissionService, domain.AttachmentCreate)).Post("/project/{project_id}/attachment", r.AttachmentController.CreateForProject)
	router.With(RequirePermission(r.PermissionService, domain.AttachmentRead)).Get("/attachment/{attachment_id}/download", r.AttachmentController.Download)
	router.With(RequirePermission(r.PermissionService, domain.AttachmentDelete)).Delete("/attachment/{attachment_id}", r.AttachmentController.Destroy)

	router.With(RequirePermission(r.PermissionService, domain.IssueRead)).Get("/issue/{issue_id}/audit-log", r.IssueAuditLogController.Index)

	router.With(RequirePermission(r.PermissionService, domain.CommentRead)).Get("/issue/{issue_id}/comment", r.CommentController.IndexByIssue)
	router.With(RequirePermission(r.PermissionService, domain.CommentCreate)).Post("/issue/{issue_id}/comment", r.CommentController.Create)
	router.With(RequirePermission(r.PermissionService, domain.CommentUpdate)).Put("/comment/{comment_id}", r.CommentController.Update)
	router.With(RequirePermission(r.PermissionService, domain.CommentDelete)).Delete("/comment/{comment_id}", r.CommentController.Destroy)

	router.With(RequirePermission(r.PermissionService, domain.IssueUpdate)).Delete("/related-issue/{related_issue_id}", r.RelatedIssueController.Delete)

	router.With(RequirePermission(r.PermissionService, domain.FilterRead)).Get("/filter/condition-options", r.FilterController.Options)
	router.With(RequirePermission(r.PermissionService, domain.FilterRead)).Get("/filter", r.FilterController.Index)
	router.With(RequirePermission(r.PermissionService, domain.FilterCreate)).Post("/filter", r.FilterController.Create)
	router.With(RequirePermission(r.PermissionService, domain.FilterRead)).Get("/filter/{filter_id}", r.FilterController.Show)
	router.With(RequirePermission(r.PermissionService, domain.FilterUpdate)).Put("/filter/{filter_id}", r.FilterController.Update)
	router.With(RequirePermission(r.PermissionService, domain.FilterDelete)).Delete("/filter/{filter_id}", r.FilterController.Destroy)

	router.With(RequirePermission(r.PermissionService, domain.BoardRead)).Get("/board", r.BoardController.Index)
	router.With(RequirePermission(r.PermissionService, domain.BoardCreate)).Post("/board", r.BoardController.Create)
	router.With(RequirePermission(r.PermissionService, domain.BoardRead)).Get("/board/{board_id}", r.BoardController.Show)
	router.With(RequirePermission(r.PermissionService, domain.BoardUpdate)).Put("/board/{board_id}", r.BoardController.Update)
	router.With(RequirePermission(r.PermissionService, domain.BoardDelete)).Delete("/board/{board_id}", r.BoardController.Destroy)

	router.With(RequirePermission(r.PermissionService, domain.DashboardRead)).Get("/dashboard", r.DashboardController.Index)
	router.With(RequirePermission(r.PermissionService, domain.DashboardCreate)).Post("/dashboard", r.DashboardController.Create)
	router.With(RequirePermission(r.PermissionService, domain.DashboardRead)).Get("/dashboard/{dashboard_id}", r.DashboardController.Show)
	router.With(RequirePermission(r.PermissionService, domain.DashboardUpdate)).Put("/dashboard/{dashboard_id}", r.DashboardController.Update)
	router.With(RequirePermission(r.PermissionService, domain.DashboardDelete)).Delete("/dashboard/{dashboard_id}", r.DashboardController.Destroy)

	router.With(RequirePermission(r.PermissionService, domain.DashboardRead)).Get("/dashboard/{dashboard_id}/widget", r.DashboardWidgetController.Index)
	router.With(RequirePermission(r.PermissionService, domain.DashboardUpdate)).Post("/dashboard/{dashboard_id}/widget", r.DashboardWidgetController.Create)
	router.With(RequirePermission(r.PermissionService, domain.DashboardUpdate)).Put("/dashboard/{dashboard_id}/widget/reorder", r.DashboardWidgetController.Reorder)
	router.With(RequirePermission(r.PermissionService, domain.DashboardUpdate)).Patch("/dashboard/{dashboard_id}/widget/{widget_id}", r.DashboardWidgetController.Update)
	router.With(RequirePermission(r.PermissionService, domain.DashboardUpdate)).Delete("/dashboard/{dashboard_id}/widget/{widget_id}", r.DashboardWidgetController.Destroy)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	router.Get("/config", ConfigController{}.Show)
	router.Post("/onboard", r.OnboardController.Create)

	router.Get("/user/permission", r.UserPermissionController.Index)

	// /user/current – saját profil, nem kell külön permission (JWT elég)
	router.Get("/user/current", r.UserCurrentController.Show)
	router.Put("/user/current", r.UserCurrentController.Update)
	router.Delete("/user/current", r.UserCurrentController.Destroy)
	router.Post("/user/register", r.UserRegisterController.Create)
	router.Post("/user/session", r.UserSessionController.Create)

	router.With(RequirePermission(r.PermissionService, domain.UserRead)).Get("/user/{user_id}/role", r.UserRoleController.Index)
	router.With(RequirePermission(r.PermissionService, domain.UserUpdate)).Post("/user/{user_id}/role", r.UserRoleController.Create)
	router.With(RequirePermission(r.PermissionService, domain.UserUpdate)).Delete("/user/{user_id}/role/{role}", r.UserRoleController.Destroy)

	router.With(RequirePermission(r.PermissionService, domain.UserRead)).Get("/user", r.UserController.Index)
	router.With(RequirePermission(r.PermissionService, domain.UserCreate)).Post("/user", r.UserController.Create)
	router.With(RequirePermission(r.PermissionService, domain.UserRead)).Get("/user/{user_id}", r.UserController.Show)
	router.With(RequirePermission(r.PermissionService, domain.UserUpdate)).Put("/user/{user_id}", r.UserController.Update)
	router.With(RequirePermission(r.PermissionService, domain.UserDelete)).Delete("/user/{user_id}", r.UserController.Destroy)

	router.With(RequirePermission(r.PermissionService, domain.TrackerRead)).Get("/tracker", r.TrackerController.Index)
	router.With(RequirePermission(r.PermissionService, domain.TrackerCreate)).Post("/tracker", r.TrackerController.Create)
	router.With(RequirePermission(r.PermissionService, domain.TrackerRead)).Get("/tracker/{tracker_id}", r.TrackerController.Show)
	router.With(RequirePermission(r.PermissionService, domain.TrackerUpdate)).Put("/tracker/{tracker_id}", r.TrackerController.Update)
	router.With(RequirePermission(r.PermissionService, domain.TrackerDelete)).Delete("/tracker/{tracker_id}", r.TrackerController.Destroy)

	router.With(RequirePermission(r.PermissionService, domain.StatusRead)).Get("/status", r.StatusController.Index)
	router.With(RequirePermission(r.PermissionService, domain.StatusCreate)).Post("/status", r.StatusController.Create)
	router.With(RequirePermission(r.PermissionService, domain.StatusRead)).Get("/status/{status_id}", r.StatusController.Show)
	router.With(RequirePermission(r.PermissionService, domain.StatusUpdate)).Put("/status/{status_id}", r.StatusController.Update)
	router.With(RequirePermission(r.PermissionService, domain.StatusDelete)).Delete("/status/{status_id}", r.StatusController.Destroy)

	router.With(RequirePermission(r.PermissionService, domain.PriorityRead)).Get("/priority", r.PriorityController.Index)
	router.With(RequirePermission(r.PermissionService, domain.PriorityCreate)).Post("/priority", r.PriorityController.Create)
	router.With(RequirePermission(r.PermissionService, domain.PriorityRead)).Get("/priority/{priority_id}", r.PriorityController.Show)
	router.With(RequirePermission(r.PermissionService, domain.PriorityUpdate)).Put("/priority/{priority_id}", r.PriorityController.Update)
	router.With(RequirePermission(r.PermissionService, domain.PriorityDelete)).Delete("/priority/{priority_id}", r.PriorityController.Destroy)

	router.With(RequirePermission(r.PermissionService, domain.TransitionRead)).Get("/transition", r.TransitionController.Index)
	router.With(RequirePermission(r.PermissionService, domain.TransitionCreate)).Post("/transition", r.TransitionController.Create)
	router.With(RequirePermission(r.PermissionService, domain.TransitionRead)).Get("/transition/{transition_id}", r.TransitionController.Show)
	router.With(RequirePermission(r.PermissionService, domain.TransitionUpdate)).Put("/transition/{transition_id}", r.TransitionController.Update)
	router.With(RequirePermission(r.PermissionService, domain.TransitionDelete)).Delete("/transition/{transition_id}", r.TransitionController.Destroy)

	router.Get("/setting", r.SettingController.Index)
	router.With(RequirePermission(r.PermissionService, domain.SettingUpdate)).Put("/setting/{key}", r.SettingController.Update)

	router.With(RequirePermission(r.PermissionService, domain.NotificationRead)).Get("/notification", r.NotificationController.Index)
	router.With(RequirePermission(r.PermissionService, domain.NotificationUpdate)).Patch("/notification/{notification_id}/read", r.NotificationController.MarkRead)

	router.Get("/ws", r.WebSocketController.Connect)

	router.With(RequirePermission(r.PermissionService, domain.ReleaseRead)).Get("/release", r.ReleaseController.Index)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseCreate)).Post("/release", r.ReleaseController.Create)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseRead)).Get("/release/{release_id}", r.ReleaseController.Show)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseUpdate)).Put("/release/{release_id}", r.ReleaseController.Update)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseDelete)).Delete("/release/{release_id}", r.ReleaseController.Destroy)

	router.With(RequirePermission(r.PermissionService, domain.ReleaseRead)).Get("/project-release", r.ProjectReleaseController.Index)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseCreate)).Post("/project-release", r.ProjectReleaseController.Create)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseRead)).Get("/project-release/{project_release_id}", r.ProjectReleaseController.Show)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseDelete)).Delete("/project-release/{project_release_id}", r.ProjectReleaseController.Destroy)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseRead)).Get("/release/{release_id}/project", r.ProjectReleaseController.IndexByRelease)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseRead)).Get("/project/{project_id}/release", r.ProjectReleaseController.IndexByProject)

	router.With(RequirePermission(r.PermissionService, domain.ReleaseRead)).Get("/project-release-issue", r.ProjectReleaseIssueController.Index)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseCreate)).Post("/project-release-issue", r.ProjectReleaseIssueController.Create)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseRead)).Get("/project-release-issue/{project_release_issue_id}", r.ProjectReleaseIssueController.Show)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseDelete)).Delete("/project-release-issue/{project_release_issue_id}", r.ProjectReleaseIssueController.Destroy)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseRead)).Get("/project-release/{project_release_id}/issue", r.ProjectReleaseIssueController.IndexByProjectRelease)
	router.With(RequirePermission(r.PermissionService, domain.ReleaseRead)).Get("/issue/{issue_id}/project-release-issue", r.ProjectReleaseIssueController.ShowByIssue)

	return router
}
