package service

import (
	"context"
	"ddd/internal/application/service/auth"
	"ddd/internal/domain/user"
	"errors"

	"github.com/google/uuid"
)

// 注册服务
type LoginService struct {
	repo         user.Repository    // domain.Repository 抽象存储对象
	TokenService *auth.TokenService // token 服务
	hasher       user.PasswordHasher
}

func NewLoginService(
	repo user.Repository,
	tokenService *auth.TokenService,
	hasher user.PasswordHasher,
) *LoginService {
	return &LoginService{
		repo:         repo,
		TokenService: tokenService,
		hasher:       hasher,
	}
}

// 登录接口
func (s *LoginService) Login(ctx context.Context, username, password string) (string, *user.User, error) {

	// 从数据库中读取 用户 domain 抽象 接口 由 infra 实现
	u, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", nil, err
	}

	// domain 密码校验
	// if err := u.CheckPassword(s.hashFn, password); err != nil {
	// 	return "", nil, err
	// }
	err = s.hasher.Compare(u.PasswordHash, password)
	if err != nil {
		return "", nil, err
	}
	token, err := s.TokenService.Generate(u.ID)
	return token, u, err
}

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type RegisterService struct {
	repo   user.Repository
	hasher user.PasswordHasher
	policy user.PasswordPolicy
}

func NewRegisterService(
	repo user.Repository,
	hasher user.PasswordHasher,
	policy user.PasswordPolicy,
) *RegisterService {
	return &RegisterService{
		repo:   repo,
		hasher: hasher,
		policy: policy,
	}
}

func (s *RegisterService) Register(
	username, password string,
) (*user.User, error) {

	// 1. 是否已存在
	if _, err := s.repo.FindByUsername(username); err == nil {
		return nil, ErrUserAlreadyExists
	}

	// 2. 创建加密后的密码
	hashed, err := s.hasher.Hash(password)

	// 2. 创建领域对象
	u, err := user.NewUser(
		uuid.NewString(),
		username,
		hashed,
		s.policy,
	)
	if err != nil {
		return nil, err
	}

	// 3. 持久化
	if err := s.repo.Save(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *LoginService) GetByID(ctx context.Context,
	id string) (*user.User, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user.User{
		ID:       u.ID,
		Username: u.Username,
	}, nil
}
