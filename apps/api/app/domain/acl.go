package domain

type Permission string

// Issue
const (
	IssueCreate Permission = "issue:create"
	IssueRead   Permission = "issue:read"
	IssueUpdate Permission = "issue:update"
	IssueDelete Permission = "issue:delete"
)

// Project
const (
	ProjectCreate Permission = "project:create"
	ProjectRead   Permission = "project:read"
	ProjectUpdate Permission = "project:update"
	ProjectDelete Permission = "project:delete"
)

// Comment
const (
	CommentCreate Permission = "comment:create"
	CommentRead   Permission = "comment:read"
	CommentUpdate Permission = "comment:update"
	CommentDelete Permission = "comment:delete"
)

// Worklog
const (
	WorklogCreate Permission = "worklog:create"
	WorklogRead   Permission = "worklog:read"
	WorklogUpdate Permission = "worklog:update"
	WorklogDelete Permission = "worklog:delete"
)

// Attachment
const (
	AttachmentCreate Permission = "attachment:create"
	AttachmentRead   Permission = "attachment:read"
	AttachmentDelete Permission = "attachment:delete"
)

// Board
const (
	BoardCreate Permission = "board:create"
	BoardRead   Permission = "board:read"
	BoardUpdate Permission = "board:update"
	BoardDelete Permission = "board:delete"
)

// Filter
const (
	FilterCreate Permission = "filter:create"
	FilterRead   Permission = "filter:read"
	FilterUpdate Permission = "filter:update"
	FilterDelete Permission = "filter:delete"
)

// Dashboard
const (
	DashboardCreate Permission = "dashboard:create"
	DashboardRead   Permission = "dashboard:read"
	DashboardUpdate Permission = "dashboard:update"
	DashboardDelete Permission = "dashboard:delete"
)

// Release
const (
	ReleaseCreate Permission = "release:create"
	ReleaseRead   Permission = "release:read"
	ReleaseUpdate Permission = "release:update"
	ReleaseDelete Permission = "release:delete"
)

// Notification
const (
	NotificationRead   Permission = "notification:read"
	NotificationUpdate Permission = "notification:update"
)

// Admin-only resources
const (
	StatusCreate   Permission = "status:create"
	StatusRead     Permission = "status:read"
	StatusUpdate   Permission = "status:update"
	StatusDelete   Permission = "status:delete"
	TrackerCreate  Permission = "tracker:create"
	TrackerRead    Permission = "tracker:read"
	TrackerUpdate  Permission = "tracker:update"
	TrackerDelete  Permission = "tracker:delete"
	PriorityCreate Permission = "priority:create"
	PriorityRead   Permission = "priority:read"
	PriorityUpdate Permission = "priority:update"
	PriorityDelete Permission = "priority:delete"
	TransitionCreate Permission = "transition:create"
	TransitionRead   Permission = "transition:read"
	TransitionUpdate Permission = "transition:update"
	TransitionDelete Permission = "transition:delete"
	UserCreate     Permission = "user:create"
	UserRead       Permission = "user:read"
	UserUpdate     Permission = "user:update"
	UserDelete     Permission = "user:delete"
	SettingRead    Permission = "setting:read"
	SettingUpdate  Permission = "setting:update"
)

type RoleDefinition struct {
	Permissions []Permission
}

var Roles = map[string]RoleDefinition{
	"admin": {
		Permissions: []Permission{
			IssueCreate, IssueRead, IssueUpdate, IssueDelete,
			ProjectCreate, ProjectRead, ProjectUpdate, ProjectDelete,
			CommentCreate, CommentRead, CommentUpdate, CommentDelete,
			WorklogCreate, WorklogRead, WorklogUpdate, WorklogDelete,
			AttachmentCreate, AttachmentRead, AttachmentDelete,
			BoardCreate, BoardRead, BoardUpdate, BoardDelete,
			FilterCreate, FilterRead, FilterUpdate, FilterDelete,
			DashboardCreate, DashboardRead, DashboardUpdate, DashboardDelete,
			ReleaseCreate, ReleaseRead, ReleaseUpdate, ReleaseDelete,
			NotificationRead, NotificationUpdate,
			StatusCreate, StatusRead, StatusUpdate, StatusDelete,
			TrackerCreate, TrackerRead, TrackerUpdate, TrackerDelete,
			PriorityCreate, PriorityRead, PriorityUpdate, PriorityDelete,
			TransitionCreate, TransitionRead, TransitionUpdate, TransitionDelete,
			UserCreate, UserRead, UserUpdate, UserDelete,
			SettingRead, SettingUpdate,
		},
	},
	"project_manager": {
		Permissions: []Permission{
			IssueCreate, IssueRead, IssueUpdate, IssueDelete,
			ProjectCreate, ProjectRead, ProjectUpdate,
			CommentCreate, CommentRead, CommentUpdate, CommentDelete,
			WorklogCreate, WorklogRead, WorklogUpdate, WorklogDelete,
			AttachmentCreate, AttachmentRead, AttachmentDelete,
			BoardRead, BoardUpdate,
			FilterCreate, FilterRead, FilterUpdate, FilterDelete,
			DashboardCreate, DashboardRead, DashboardUpdate, DashboardDelete,
			ReleaseCreate, ReleaseRead, ReleaseUpdate, ReleaseDelete,
			NotificationRead, NotificationUpdate,
			StatusRead, TrackerRead, PriorityRead, TransitionRead,
			UserRead,
			SettingRead,
		},
	},
	"developer": {
		Permissions: []Permission{
			IssueCreate, IssueRead, IssueUpdate,
			ProjectRead,
			CommentCreate, CommentRead, CommentUpdate, CommentDelete,
			WorklogCreate, WorklogRead, WorklogUpdate, WorklogDelete,
			AttachmentCreate, AttachmentRead, AttachmentDelete,
			BoardRead,
			FilterCreate, FilterRead, FilterUpdate, FilterDelete,
			DashboardCreate, DashboardRead, DashboardUpdate, DashboardDelete,
			ReleaseRead,
			NotificationRead, NotificationUpdate,
			StatusRead, TrackerRead, PriorityRead, TransitionRead,
		},
	},
	"viewer": {
		Permissions: []Permission{
			IssueRead,
			ProjectRead,
			CommentRead,
			WorklogRead,
			AttachmentRead,
			BoardRead,
			FilterRead,
			DashboardRead,
			ReleaseRead,
			NotificationRead,
			StatusRead, TrackerRead, PriorityRead, TransitionRead,
		},
	},
}

func RoleHasPermission(roleName string, perm Permission) bool {
	role, ok := Roles[roleName]
	if !ok {
		return false
	}
	for _, p := range role.Permissions {
		if p == perm {
			return true
		}
	}
	return false
}
