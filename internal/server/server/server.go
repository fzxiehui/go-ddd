package server

import (
	"context"
	"ddd/internal/interface/grpc"
	"log"
	"net/http"
)

type Server struct {
	HTTP *http.Server
	GRPC *grpc.GRPCServer
}

func NewServer(httpSrv *http.Server, grpcSrv *grpc.GRPCServer) *Server {
	return &Server{
		HTTP: httpSrv,
		GRPC: grpcSrv,
	}
}

func (s *Server) Run() error {
	log.Println("http server started")
	return s.HTTP.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("shutting down server...")

	if err := s.HTTP.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("server shutdown complete")
	return nil
}

func (s *Server) RunGrpc() error {
	log.Println("grpc server started")
	return s.GRPC.Start(":9090")
}
