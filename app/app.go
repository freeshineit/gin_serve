package app

import (
	"gin_serve/app/config"
	"gin_serve/app/middleware"
	"gin_serve/app/routes"
	"gin_serve/swagger"
	"gin_serve/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunServer(conf config.ServerConfig) error {

	if conf.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// viper /app/config/config.yaml
	// read /cmd/app.go

	// hooks, config,...

	// 连接mysql数据库
	config.SetupDatabaseConnection()

	// 连接redis
	config.SetupRedisConnection()

	// r := gin.Default()
	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())

	// 中间件
	middleware.SetMiddleware(r)

	// api docs
	swagger.InitSwagger(r)

	// set up routes
	routes.SetupRoutes(r)

	srv := &http.Server{
		Addr:    ":" + conf.Port,
		Handler: r,
	}

	log.Printf("listen: \033[1;32;40mhttp://localhost:%s\033[0m\n", conf.Port)

	err := utils.ListenAndServe(srv)

	if err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

	return err
}
