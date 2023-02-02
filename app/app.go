package app

import (
	"gin_serve/app/config"
	"gin_serve/app/middleware"
	"gin_serve/app/model"
	"gin_serve/app/routes"
	"gin_serve/helpers"
	"gin_serve/swagger"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "gin_serve/app/utils"
)

func RunServer(conf config.ServerConfig) error {

	if conf.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	config.SetupDatabaseConnection()
	config.SetupRedisConnection()

	model.GormAutoMigration(config.DB)

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

	log.Printf("listen: %s\n", helpers.ColorBlueString("http://localhost:"+conf.Port))

	err := helpers.ListenAndServe(srv)

	if err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

	// close
	defer config.CloseMysqlConnection(config.DB)
	defer config.CloseRedisConnection(config.RedisClient)

	return err
}
