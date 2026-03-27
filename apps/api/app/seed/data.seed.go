package seed

type SeedData struct {
	Trackers       []TrackerData       `yaml:"trackers" validate:"required,dive"`
	Statuses       []StatusData        `yaml:"statuses" validate:"required,dive"`
	Priorities     []PriorityData      `yaml:"priorities" validate:"dive"`
	Boards         []BoardData         `yaml:"boards" validate:"required,dive"`
	Users          []UserData          `yaml:"users" validate:"required,dive"`
	Dashboards       []DashboardData       `yaml:"dashboards" validate:"required,dive"`
	DashboardWidgets []DashboardWidgetData `yaml:"dashboard_widgets" validate:"dive"`
	Filters          []FilterData          `yaml:"filters" validate:"required,dive"`
	Projects       []ProjectData       `yaml:"projects" validate:"required,dive"`
	Issues         []IssueData         `yaml:"issues" validate:"required,dive"`
	RelatedIssues  []RelatedIssueData  `yaml:"related_issues"`
	Notifications  []NotificationData  `yaml:"notifications" validate:"dive"`
}
