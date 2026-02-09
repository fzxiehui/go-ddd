package server

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	HTTP *http.Server
	// GRPC *grpc.Server   // TODO
}

func NewServer(httpSrv *http.Server) *Server {
	return &Server{
		HTTP: httpSrv,
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
