package seed

type SeedData struct {
	Trackers       []TrackerData       `yaml:"trackers" validate:"required,dive"`
	Statuses       []StatusData        `yaml:"statuses" validate:"required,dive"`
	Boards         []BoardData         `yaml:"boards" validate:"required,dive"`
	Users          []UserData          `yaml:"users" validate:"required,dive"`
	Dashboards     []DashboardData     `yaml:"dashboards" validate:"required,dive"`
	Filters        []FilterData        `yaml:"filters" validate:"required,dive"`
	Projects       []ProjectData       `yaml:"projects" validate:"required,dive"`
	Issues         []IssueData         `yaml:"issues" validate:"required,dive"`
	RelatedIssues  []RelatedIssueData  `yaml:"related_issues"`
}
