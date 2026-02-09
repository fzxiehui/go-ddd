package server

import "github.com/gin-gonic/gin"

type Server struct {
	HTTP *gin.Engine
	// GRPC *grpc.Server   // TODO
}

func NewServer(http *gin.Engine) *Server {
	return &Server{
		HTTP: http,
	}
}

func (s *Server) Run() error {
	return s.HTTP.Run(":8080")
}
