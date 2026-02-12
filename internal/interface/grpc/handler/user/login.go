package user

import (
	"context"

	userv1 "ddd/api/gen/user/v1"
	user "ddd/internal/application/service/user"
)

type AuthHandler struct {
	userv1.UnimplementedAuthServiceServer
	loginService *user.LoginService
}

func NewAuthHandler(loginService *user.LoginService) *AuthHandler {
	return &AuthHandler{
		loginService: loginService,
	}
}

func (h *AuthHandler) Login(
	ctx context.Context,
	req *userv1.LoginRequest,
) (*userv1.LoginReply, error) {

	token, u, err := h.loginService.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &userv1.LoginReply{
		UserId: u.ID,
		Token:  token,
	}, nil
}
