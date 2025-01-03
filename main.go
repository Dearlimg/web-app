package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web-app/controllers"
	"web-app/dao/mysql"
	"web-app/dao/redis"
	"web-app/logger"
	"web-app/pkg/snowflake"
	"web-app/routes"
	"web-app/settings"

	"github.com/spf13/viper"

	"go.uber.org/zap"

	_ "web-app/docs" // 千万不要忘了导入把你上一步生成的docs
)

// @title bluebell
// @version 1.0
// @description 简单贴吧
// @termsOfService http://swagger.io/terms/

// @contact.name 1492568061@qq.com
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath ./swag

func main() {
	//1.加载配置
	for i := 0; i < 3; i++ {
		fmt.Println(1)
	}
	if err := settings.Init(); err != nil {
		fmt.Println("init settings failed:", err)
		return
	}
	//2.加载日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Println("init logger failed:", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")
	//3.处理mysql
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Println("init mysql failed:", err)
		return
	}
	defer mysql.Close()
	//4.redis
	if err := redis.Init(); err != nil {
		fmt.Println("init redis failed:", err)
		return
	}
	defer redis.Close()

	//fmt.Println(settings.Conf.StartTime, settings.Conf.MachineID)
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Println("init snowflake failed:", err)
		return
	}

	if err := controllers.Init("zh"); err != nil {
		fmt.Println("init translator failed:", err)
		return
	}

	//5.注册路由
	r := routes.Init(settings.Conf.Mode)
	//6.启动服务(优雅关机,平滑启动)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
