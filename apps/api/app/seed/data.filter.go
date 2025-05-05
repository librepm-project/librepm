package seed

type FilterData struct {
	Name        string          `yaml:"name"`
	Description string          `yaml:"description"`
	Conditions  []ConditionData `yaml:"conditions"`
	UserEmail   string          `yaml:"user_email"`
}

type ConditionData struct {
	Condition string `yaml:"condition"`
}
