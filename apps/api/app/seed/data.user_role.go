package seed

type UserRoleData struct {
	UserEmail string `yaml:"user_email" validate:"required,email"`
	Role      string `yaml:"role" validate:"required"`
}
