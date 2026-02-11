package userinfra

import (
	"time"

	"gorm.io/gorm"
)

type UserPO struct {
	ID           string `gorm:"primaryKey"`
	Username     string `gorm:"uniqueIndex"`
	PasswordHash string
	CreatedAt    time.Time
}

func (u *UserPO) TableName() string {
	return "users"
}

type PasswordPolicyPO struct {
	gorm.Model
	MinLength     int  // 最短长度
	RequireNumber bool // 是否一定要包含数字
}

func (u *PasswordPolicyPO) TableName() string {
	return "password_policy"
}
