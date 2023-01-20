package app

import (
	"gin_server/app/config"
	"gin_server/app/middleware"
	"gin_server/app/routes"
	"gin_server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RunServer(conf config.ServerConfig) error {

	if conf.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// viper /app/config/config.yaml
	// read /cmd/app.go

	log := logrus.New()
	// hooks, config,...

	// 连接mysql数据库
	config.SetupDatabaseConnection()

	// 连接redis
	config.SetupRedisConnection()

	// r := gin.Default()
	r := gin.New()
	r.Use(middleware.Logger(log), gin.Recovery())

	// 中间件
	middleware.SetMiddleware(r)

	// set up routes
	routes.SetupRoutes(r)

	srv := &http.Server{
		Addr:    ":" + conf.Port,
		Handler: r,
	}

	log.Printf("listen: http://localhost:%s\n", conf.Port)

	err := utils.ListenAndServe(srv)

	if err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

	return err
}
