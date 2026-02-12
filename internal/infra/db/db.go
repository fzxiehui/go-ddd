package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ddd/internal/config"
	"ddd/internal/infra/security"
	userinfra "ddd/internal/infra/user"
)

func InitSQLite(cfg *config.Config,
	bph *security.BcryptPasswordHasher) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DB.Name), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 统一迁移
	if err := db.AutoMigrate(
		&userinfra.UserPO{},
		&userinfra.PasswordPolicyPO{},
	); err != nil {
		return nil, err
	}

	// 初始化默认数据
	// 默认密码规则
	if err := userinfra.EnsureDefaultPasswordPolicy(db); err != nil {
		return nil, err
	}

	// root 用户
	if err := userinfra.EnsureDefaultUser(db, bph); err != nil {
		return nil, err
	}

	return db, nil
}
