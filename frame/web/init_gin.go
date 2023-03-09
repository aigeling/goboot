package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mattn/go-colorable"
)

type GinConfig struct {
	//是否debug模式启动
	Debug bool

	//启动端口
	Port int
	//是否允许跨域
	AllowCors bool
	//允许跨域的域名列表
	AllowOrigins []string

	ValidFuncs map[string]func(fl validator.FieldLevel) bool

	//启动session管理
	// EnableSession bool
	// //session存储
	// SessionStore *sessions.Store

	// //session存活秒数
	// SessionMaxAge int
}

// 启动gin服务
func StartGinServer(initRouter func(ge *gin.Engine), conf *GinConfig) {
	if conf.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DefaultWriter = colorable.NewColorableStdout()

	router := gin.New()

	router.Use(ginLogger, ginRecovery(true))
	// router := gin.Default()

	if conf.AllowCors {
		initCors(router, conf)
	}

	// if conf.EnableSession {
	// 	initSession(router, conf.SessionMaxAge, *conf.SessionStore)
	// }

	initRouter(router)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: router,
	}

	if conf.Debug {

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	} else {
		go func() {
			// 开启一个goroutine启动服务
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		log.Printf("Gin [%d]Start ... ", conf.Port)
		// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
		quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
		// kill 默认会发送 syscall.SIGTERM 信号
		// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
		// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
		// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
		<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
		log.Println("Shutdown Server ...")
		// 创建一个5秒超时的context
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown: ", err)
		}

	}
	log.Println("Server exiting")

}

func initCors(ge *gin.Engine, conf *GinConfig) {
	config := cors.DefaultConfig()
	config.AllowOrigins = conf.AllowOrigins //[]string{"http://127.0.0.1:8080", "http://localhost:8080"}
	config.AllowCredentials = true
	ge.Use(cors.New(config))
}
