package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	userinfra "ddd/internal/infra/user"
)

type DBConfig struct {
	Path string
}

func InitSQLite(opt DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(opt.Path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 统一做迁移（infra 层允许）
	if err := db.AutoMigrate(
		&userinfra.UserPO{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
