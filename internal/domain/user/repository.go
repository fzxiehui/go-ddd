package user

import "context"

type Repository interface {
	FindByUsername(username string) (*User, error)
	Save(u *User) error
	FindByID(ctx context.Context, id string) (*User, error)
}
