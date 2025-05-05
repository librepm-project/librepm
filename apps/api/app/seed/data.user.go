package seed

type UserData struct {
	Email        string `yaml:"email"`
	PasswordHash string `yaml:"password_hash"`
	FirstName    string `yaml:"first_name"`
	LastName     string `yaml:"last_name"`
	Phone        string `yaml:"phone"`
	Language     string `yaml:"language"`
	Country      string `yaml:"country"`
	Blocked      bool   `yaml:"blocked"`
}
