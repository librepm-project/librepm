package seed

type UserData struct {
	Email        string `yaml:"email" validate:"required"`
	PasswordHash string `yaml:"password_hash" validate:"required"`
	FirstName    string `yaml:"first_name" validate:"required"`
	LastName     string `yaml:"last_name" validate:"required"`
	Phone        string `yaml:"phone" validate:"required"`
	Language     string `yaml:"language" validate:"required"`
	Country      string `yaml:"country" validate:"required"`
	Blocked      bool   `yaml:"blocked"`
}
