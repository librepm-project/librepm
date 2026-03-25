package seed

type BoardData struct {
	Name        string            `yaml:"name" validate:"required"`
	Description string            `yaml:"description" validate:"required"`
	Columns     []BoardColumnData `yaml:"columns" validate:"required,dive"`
}

type BoardColumnData struct {
	Label    string   `yaml:"label" validate:"required"`
	Statuses []string `yaml:"statuses"`
}
