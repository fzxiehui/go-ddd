package user

import (
	"errors"
)

type User struct {
	ID           string
	Username     string
	PasswordHash string
}

// var (
// 	ErrInvalidPassword = errors.New("invalid password")
// )

// func (u *User) CheckPassword(hashFn func(string) string, password string) error {
// 	if u.PasswordHash != hashFn(password) {
// 		return ErrInvalidPassword
// 	}
// 	return nil
// }

var (
	ErrUsernameEmpty = errors.New("username empty")
	ErrPasswordEmpty = errors.New("password empty")
)

func NewUser(
	id string,
	username string,
	rawPassword string,
	policy PasswordPolicy,
) (*User, error) {

	if err := policy.Validate(rawPassword); err != nil {
		return nil, err
	}

	if username == "" {
		return nil, ErrUsernameEmpty
	}
	if rawPassword == "" {
		return nil, ErrPasswordEmpty
	}

	return &User{
		ID:           id,
		Username:     username,
		PasswordHash: rawPassword,
	}, nil
}
