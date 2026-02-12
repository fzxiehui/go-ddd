package userinfra

import (
	"ddd/internal/infra/security"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func EnsureDefaultPasswordPolicy(db *gorm.DB) error {
	var count int64

	if err := db.Model(&PasswordPolicyPO{}).Count(&count).Error; err != nil {
		return err
	}

	// 如果没有规则，则创建默认规则
	if count == 0 {
		defaultPolicy := PasswordPolicyPO{
			MinLength:     3,
			RequireNumber: false,
		}
		if err := db.Create(&defaultPolicy).Error; err != nil {
			return err
		}
	}

	return nil
}

func EnsureDefaultUser(db *gorm.DB,
	bph *security.BcryptPasswordHasher) error {
	var count int64
	if err := db.Model(&UserPO{}).Where("username == ?", "root").Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		password, err := bph.Hash("root")
		if err != nil {
			return nil
		}
		root := UserPO{
			ID:           uuid.NewString(),
			Username:     "root",
			PasswordHash: password,
		}
		if err := db.Create(&root).Error; err != nil {
			return err
		}
	}

	return nil
}
