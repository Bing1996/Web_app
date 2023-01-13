package main

import (
	"Web_App/asset/settings"
	"Web_App/middleware/logger"
	"Web_App/repository/mysql"
	"Web_App/repository/redis"
	"Web_App/router"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// 各个服务初始化，包括数据库、Redis、配置文件、中间件等
	appInit()

	// 路由初始化
	r := router.Setup()

	// 启动服务
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal(fmt.Sprintf("listen: %s", err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	// 创建一个接收信号的通道
	quit := make(chan os.Signal, 1)
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("server shutdown...")
	// 创建一个5秒的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒后强制退出
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal(fmt.Sprintln("Server exiting ", err))
	}

	zap.L().Info("Server exiting")

}

func appInit() {
	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init setting failed, err: %s", err)
		return
	}

	// 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init log failed, err: %s", err)
		return
	}
	defer zap.L().Sync()

	// 初始化Mysql，利用GORM框架与mysql引擎进行连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err: %s", err)
		return
	}
	defer mysql.Close()

	// 初始化Redis
	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed, err: %s", err)
		return
	}
	defer redis.Close()
}
