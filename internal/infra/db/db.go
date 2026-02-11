package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ddd/internal/config"
	userinfra "ddd/internal/infra/user"
)

func InitSQLite(cfg *config.Config) (*gorm.DB, error) {
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
	if err := userinfra.EnsureDefaultPasswordPolicy(db); err != nil {
		return nil, err
	}

	return db, nil
}
