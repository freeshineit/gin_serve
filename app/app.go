package app

import (
	"context"
	"errors"
	"go_python_serve/app/middleware"
	"go_python_serve/app/models"
	"go_python_serve/app/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ServerConfig struct {
	Port string
	Mode string
}

func RunServer(config ServerConfig) {

	if config.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// viper config.yaml
	// config.SetConfig()

	log := logrus.New()
	// hooks, config,...

	// 连接mysql数据库
	models.InitMysqlDB()

	// 连接redis
	models.InitRedis()

	// r := gin.Default()
	r := gin.New()
	r.Use(middleware.Logger(log), gin.Recovery())

	// 中间件
	middleware.SetMiddleware(r)

	// set up routes
	routes.SetupRoutes(r)

	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: r,
	}

	log.Printf("listen: http://localhost:%s\n", config.Port)
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
