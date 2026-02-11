package service

import (
	"ddd/internal/application/service/auth"
	"ddd/internal/domain/user"
	"errors"

	"github.com/google/uuid"
)

// 注册服务
type LoginService struct {
	repo         user.Repository     // domain.Repository 抽象存储对象
	TokenService *auth.TokenService  // token 服务
	hashFn       func(string) string // 密码生成函数
}

func NewLoginService(
	repo user.Repository,
	tokenService *auth.TokenService,
	hashFn func(string) string,
) *LoginService {
	return &LoginService{
		repo:         repo,
		hashFn:       hashFn,
		TokenService: tokenService,
	}
}

// 登录接口
func (s *LoginService) Login(username, password string) (string, *user.User, error) {

	// 从数据库中读取 用户 domain 抽象 接口 由 infra 实现
	u, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", nil, err
	}

	// domain 密码校验
	if err := u.CheckPassword(s.hashFn, password); err != nil {
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
	hashFn func(string) string
	policy user.PasswordPolicy
}

func NewRegisterService(
	repo user.Repository,
	hashFn func(string) string,
	policy user.PasswordPolicy,
) *RegisterService {
	return &RegisterService{
		repo:   repo,
		hashFn: hashFn,
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

	// 2. 创建领域对象
	u, err := user.NewUser(
		uuid.NewString(),
		username,
		password,
		s.hashFn,
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
