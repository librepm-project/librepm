package seed

type BoardData struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Columns     []BoardColumnData `yaml:"columns"`
}

type BoardColumnData struct {
	Label string `yaml:"label"`
}
