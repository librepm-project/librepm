package seed

type TrackerData struct {
	Name  string `yaml:"name" validate:"required"`
	Color string `yaml:"color"`
}
