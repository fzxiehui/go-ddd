package userinfra

import (
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
		return db.Create(&defaultPolicy).Error
	}

	return nil
}
