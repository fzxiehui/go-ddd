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

/*
 * 开启定时任务
 */
func (s *Server) RunScheduler() {
	log.Println("scheduler started")
	s.Scheduler.Start()
	// 定时任务 TODO: 后续加到 handler中
	// s.Scheduler.AddJob(
	// 	"print_time",
	// 	"*/5 * * * * *", // 每5秒
	// 	func() {
	// 		fmt.Println("执行定时任务")
	// 	},
	// )
}

/*
 * 退出定时任务
 */
func (s *Server) ShutdownScheduler() {
	s.Scheduler.Stop()
	log.Println("scheduler shutdown complete")
}

/*
 * 启动HTTP服务 (阻塞运行)
 */
func (s *Server) RunHTTP() error {
	log.Println("http server started")
	return s.HTTP.ListenAndServe()
}

/*
 * 退出http服务
 */
func (s *Server) ShutdownHTTP(ctx context.Context) error {
	log.Println("shutting down server...")

	if err := s.HTTP.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("http server shutdown complete")
	return nil
}

/*
 * 启动grpc服务 (阻塞运行)
 */
func (s *Server) RunGrpc() error {
	log.Println("grpc server started")
	return s.GRPC.Start(":9090")
}

/*
 * 退出grpc服务
 */
func (s *Server) ShutdownGrpc() error {
	log.Println("shutting down grpc server ...")
	s.GRPC.Shutdown()
	log.Println("grpc server shutdown complete")
	return nil
}
