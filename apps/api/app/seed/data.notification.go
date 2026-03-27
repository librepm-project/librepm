package seed

type NotificationData struct {
	UserEmail  string `yaml:"user_email" validate:"required,email"`
	Type       string `yaml:"type" validate:"required"`
	EntityType string `yaml:"entity_type"`
	Read       bool   `yaml:"read"`
}
