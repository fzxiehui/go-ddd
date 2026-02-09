package user

type Repository interface {
	FindByUsername(username string) (*User, error)
	Save(u *User) error
}
