package grpc

import (
	"net"

	userv1 "ddd/api/gen/user/v1"
	"ddd/internal/interface/grpc/handler/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	server *grpc.Server
}

func NewGRPCServer(authHandler *user.AuthHandler) *GRPCServer {
	s := grpc.NewServer()

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
