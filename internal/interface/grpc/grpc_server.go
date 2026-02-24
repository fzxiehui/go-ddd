package grpc

import (
	"net"

	userv1 "ddd/api/gen/user/v1"
	"ddd/internal/application/service/auth"
	"ddd/internal/interface/grpc/handler/user"
	"ddd/internal/interface/grpc/middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	server *grpc.Server
}

func NewGRPCServer(authHandler *user.AuthHandler,
	tokenSvc *auth.TokenService) *GRPCServer {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			middleware.JWTUnaryInterceptor(tokenSvc),
		),
	)

	reflection.Register(s)

	userv1.RegisterAuthServiceServer(s, authHandler)

	return &GRPCServer{
		server: s,
	}
}

func (g *GRPCServer) Start(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return g.server.Serve(lis)
}

func (g *GRPCServer) Shutdown() {
	g.server.Stop()
}
