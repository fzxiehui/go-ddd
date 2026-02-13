package security

import (
	"ddd/internal/config"
	"ddd/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordHasher struct {
	cost int
}

func NewBcryptPasswordHasher(cfg *config.Config) user.PasswordHasher {
	return &BcryptPasswordHasher{
		cost: cfg.Security.Cost,
	}
}

func (b *BcryptPasswordHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	return string(bytes), err
}

func (b *BcryptPasswordHasher) Compare(hashed string, plain string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(plain),
	)
}
