package seed

type ProjectTrackerData struct {
	Name     string   `yaml:"name"`
	Statuses []string `yaml:"statuses"`
}

type ProjectData struct {
	Name     string               `yaml:"name"`
	CodeName string               `yaml:"code_name"`
	Trackers []ProjectTrackerData `yaml:"trackers"`
}
