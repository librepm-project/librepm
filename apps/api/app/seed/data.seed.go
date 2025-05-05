package seed

type SeedData struct {
	Trackers   []TrackerData   `yaml:"trackers"`
	Statuses   []StatusData    `yaml:"statuses"`
	Boards     []BoardData     `yaml:"boards"`
	Users      []UserData      `yaml:"users"`
	Dashboards []DashboardData `yaml:"dashboards"`
	Filters    []FilterData    `yaml:"filters"`
	Projects   []ProjectData   `yaml:"projects"`
}
