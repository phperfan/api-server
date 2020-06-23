package main

import (
	"api-server/config"
	"api-server/handler"
	"api-server/internal/model"
	"api-server/pkg/log"
	routers "api-server/router"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var cfg = pflag.StringP("config", "c", "", "goapi config file path.")

// golang api项目骨架
func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// Set gin mode.
	gin.SetMode(viper.GetString("run_mode"))

	// Create the Gin engine.
	router := gin.Default()

	// HealthCheck 健康检查路由
	router.GET("/health", handler.HealthCheck)
	// metrics router 可以在 prometheus 中进行监控
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API Routes.
	routers.Load(router)

	log.Al.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	srv := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Al.Fatalf("listen: %s", err.Error())
		}
	}()

	//schedule.Init()
	//email.Init()

	gracefulStop(srv)
}

// gracefulStop 优雅退出
// 等待中断信号以超时 5 秒正常关闭服务器
// 官方说明：https://github.com/gin-gonic/gin#graceful-restart-or-stop
func gracefulStop(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	// kill 命令发送信号 syscall.SIGTERM
	// kill -2 命令发送信号 syscall.SIGINT
	// kill -9 命令发送信号 syscall.SIGKILL
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Al.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Al.Fatal("Server Shutdown:", err)
	}
	// 5 秒后捕获 ctx.Done() 信号
	select {
	case <-ctx.Done():
		log.Al.Info("timeout of 5 seconds.")
	default:
	}
	log.Al.Info("Server exiting")
}
