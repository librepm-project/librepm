package seed

type NotificationData struct {
	UserEmail  string `yaml:"user_email" validate:"required,email"`
	Type       string `yaml:"type" validate:"required"`
	EntityType string `yaml:"entity_type"`
	IssueKey   string `yaml:"issue_key"`
	Read       bool   `yaml:"read"`
}
