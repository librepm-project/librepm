package seed

type ProjectTrackerData struct {
	Name     string   `yaml:"name" validate:"required"`
	Statuses []string `yaml:"statuses" validate:"required"`
}

type ProjectData struct {
	Name     string               `yaml:"name" validate:"required"`
	CodeName string               `yaml:"code_name" validate:"required"`
	Trackers []ProjectTrackerData `yaml:"trackers" validate:"required,dive"`
}
