package seed

type FilterData struct {
	Name        string          `yaml:"name" validate:"required"`
	Description string          `yaml:"description" validate:"required"`
	Conditions  []ConditionData `yaml:"conditions" validate:"required,dive"`
	UserEmail   string          `yaml:"user_email" validate:"required"`
}

type ConditionData struct {
	Condition string `yaml:"condition" validate:"required"`
}
