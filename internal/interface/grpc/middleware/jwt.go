package middleware

import (
	"context"
	"ddd/internal/application/service/auth"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const UserIDKey contextKey = "user_id"

type TokenService interface {
	Parse(token string) (string, error) // 返回 userID
}

func JWTUnaryInterceptor(tokenSvc *auth.TokenService) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		// 放行 Login
		if strings.Contains(info.FullMethod, "Login") {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(16, "missing metadata")
		}

		authHeader := md.Get("authorization")
		if len(authHeader) == 0 {
			return nil, status.Errorf(16, "missing token")
		}

		token := strings.TrimPrefix(authHeader[0], "Bearer ")

		claims, err := tokenSvc.Parse(token)
		if err != nil {
			return nil, status.Errorf(16, "invalid token")
		}

		ctx = context.WithValue(ctx, UserIDKey, claims.UserID)

		return handler(ctx, req)
	}
}
