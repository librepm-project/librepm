package seed

type FilterData struct {
	Name        string          `yaml:"name" validate:"required"`
	Description string          `yaml:"description" validate:"required"`
	Conditions  []ConditionData `yaml:"conditions" validate:"required,dive"`
	UserEmail   string          `yaml:"user_email" validate:"required"`
}

type ConditionData struct {
	Field string `yaml:"field" validate:"required"`
	Op    string `yaml:"op" validate:"required"`
	Value string `yaml:"value" validate:"required"`
}
