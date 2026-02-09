package userinfra

import (
	"errors"

	"gorm.io/gorm"
)

/*
 * 密码规则
 */
type DBPasswordPolicy struct {
	db *gorm.DB // grom 对象
}

func NewDBPasswordPolicy(db *gorm.DB) *DBPasswordPolicy {
	return &DBPasswordPolicy{db: db}
}

/*
 * 密码校验
 */
func (p *DBPasswordPolicy) Validate(password string) error {
	var cfg struct {
		MinLength     int  // 最短长度
		RequireNumber bool // 是否一定要包含数字
	}

	// 查询 sqlite 表
	if err := p.db.Table("password_policy").First(&cfg).Error; err != nil {
		return err
	}

	// 密码长度校验
	if len(password) < cfg.MinLength {
		return errors.New("password too short")
	}

	// 是否包含数字校验
	if cfg.RequireNumber {
		hasNum := false
		for _, c := range password {
			if c >= '0' && c <= '9' {
				hasNum = true
				break
			}
		}
		if !hasNum {
			return errors.New("password must contain number")
		}
	}

	return nil
}
