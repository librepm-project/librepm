package seed

type StatusData struct {
	Name  string `yaml:"name" validate:"required"`
	Color string `yaml:"color"`
}
