package seed

type PriorityData struct {
	Name  string `yaml:"name" validate:"required"`
	Color string `yaml:"color"`
}
