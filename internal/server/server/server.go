package server

import (
	"context"
	"ddd/internal/application/service/job"
	"ddd/internal/interface/grpc"
	"log"
	"net/http"
)

type Server struct {
	HTTP      *http.Server
	GRPC      *grpc.GRPCServer
	Scheduler *job.Scheduler
}

func NewServer(
	httpSrv *http.Server,
	grpcSrv *grpc.GRPCServer,
	scheduler *job.Scheduler) *Server {
	return &Server{
		HTTP:      httpSrv,
		GRPC:      grpcSrv,
		Scheduler: scheduler,
	}
}

func (s *Server) Run() error {
	log.Println("http server started")
	s.Scheduler.Start()
	// 定时任务 TODO: 后续加到 handler中
	// s.Scheduler.AddJob(
	// 	"print_time",
	// 	"*/5 * * * * *", // 每5秒
	// 	func() {
	// 		fmt.Println("执行定时任务")
	// 	},
	// )
	return s.HTTP.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("shutting down server...")
	s.Scheduler.Stop()

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
