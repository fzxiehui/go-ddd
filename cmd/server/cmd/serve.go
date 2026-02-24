package cmd

import (
	"context"
	"ddd/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

/*
 * 启动服务
 *	step 1 : 加载 配置文件
 *	step 2 : 启动 http 服务, GRPC 服务, 定时任务管理器
 *	step 3 : 等待 退出信息号
 *	step 4 : 退出 http 服务, GRPC 服务, 定时任务管理器
 */
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Server(http & grpc)",
	RunE: func(cmd *cobra.Command, args []string) error {

		/*
		 * step 1 : 加载 配置文件
		 */
		cfg, err := config.Load(configFile)
		app, err := InitServer(cfg)
		if err != nil {
			panic(err)
		}

		/*
		 * step 2 : 启动 http 服务, GRPC 服务, 定时任务管理器
		 */
		// 2.1 启动 http 服务
		go func() {
			if err := app.RunHTTP(); err != nil {
				log.Println("server stopped:", err)
			}
		}()

		// 2.2 启动 grpc 服务
		go func() {
			if err := app.RunGrpc(); err != nil {
				log.Println("server stopped:", err)
			}
		}()

		// 2.3 启动 定时任务管理器
		app.RunScheduler()

		/*
		 * step 3 : 等待 退出信息号
		 */
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		// 优雅关闭
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		/*
		 * step 4 : 退出 http 服务, GRPC 服务, 定时任务管理器
		 */

		// 4.1 退出 http 服务
		if err := app.ShutdownHTTP(ctx); err != nil {
			log.Println("forced shutdown:", err)
		}

		// 4.2 退出 GRPC 服务
		app.ShutdownGrpc()

		// 4.3 退出 定时任务管理器
		app.ShutdownScheduler()
		return nil

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
