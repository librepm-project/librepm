package seed

type DashboardData struct {
	Name        string `yaml:"name" validate:"required"`
	Description string `yaml:"description" validate:"required"`
	UserEmail   string `yaml:"user_email" validate:"required"`
}
