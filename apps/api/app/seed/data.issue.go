package seed

type IssueData struct {
	Summary     string `yaml:"summary" validate:"required"`
	Description string `yaml:"description" validate:"required"`
	ProjectName string `yaml:"project_name" validate:"required"`
	TrackerName string `yaml:"tracker_name" validate:"required"`
	StatusName  string `yaml:"status_name" validate:"required"`
	ParentName  string `yaml:"parent_name"`
}
