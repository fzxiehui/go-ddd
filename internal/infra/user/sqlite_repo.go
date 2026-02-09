package userinfra

import (
	"errors"

	"gorm.io/gorm"

	"ddd/internal/domain/user"
)

/*
 * 用户数据库对象
 */
type SQLiteRepo struct {
	db *gorm.DB
}

func NewSQLiteRepo(db *gorm.DB) *SQLiteRepo {
	return &SQLiteRepo{db: db}
}

/*
 * 通过用户名查找用户
 */
func (r *SQLiteRepo) FindByUsername(username string) (*user.User, error) {
	var po UserPO
	err := r.db.Where("username = ?", username).First(&po).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &user.User{
		ID:           po.ID,
		Username:     po.Username,
		PasswordHash: po.PasswordHash,
	}, nil
}

/*
 * 创建用户
 */
func (r *SQLiteRepo) Save(u *user.User) error {
	po := UserPO{
		ID:           u.ID,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}
	return r.db.Create(&po).Error
}
