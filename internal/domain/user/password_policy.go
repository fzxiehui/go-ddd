package user

type PasswordPolicy interface {
	Validate(password string) error
}
