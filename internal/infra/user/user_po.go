package userinfra

import "time"

type UserPO struct {
	ID           string `gorm:"primaryKey"`
	Username     string `gorm:"uniqueIndex"`
	PasswordHash string
	CreatedAt    time.Time
}
