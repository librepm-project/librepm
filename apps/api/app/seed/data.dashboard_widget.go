package seed

type DashboardWidgetData struct {
	DashboardName string `yaml:"dashboard_name" validate:"required"`
	FilterName    string `yaml:"filter_name"`
	Type          string `yaml:"type" validate:"required"`
	Title         string `yaml:"title" validate:"required"`
	Weight        int    `yaml:"weight"`
	Width         string `yaml:"width"`
}
