package main

import (
	"context"
	"ddd/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.Load("")
	app, err := InitServer(cfg)
	if err != nil {
		panic(err)
	}

	// 启动 server
	go func() {
		if err := app.Run(); err != nil {
			log.Println("server stopped:", err)
		}
	}()

	// 监听退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Println("forced shutdown:", err)
	}
}
