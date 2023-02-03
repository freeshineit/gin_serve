package app

import (
	"gin_serve/app/middleware"
	"gin_serve/app/model"
	"gin_serve/app/routes"
	"gin_serve/config"
	"gin_serve/helper"
	"gin_serve/swagger"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// app serve
func Serve(mode string) error {

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	port := config.Conf.App.Port

	DB := config.SetupDatabaseConnection()
	config.SetupRedisConnection()

	model.GormAutoMigration(DB)

	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())

	// 中间件
	middleware.SetMiddleware(r)

	if mode != "release" {
		// api docs
		swagger.InitSwagger(r)
	}

	// set up routes
	routes.SetupRoutes(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Printf("listen: %s\n", helper.ColorBlueString("http://localhost:"+port))

	err := helper.ListenAndServe(srv)

	if err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

	// close
	defer config.CloseMysqlConnection(DB)
	defer config.CloseRedisConnection(config.RedisClient)

	return err
}
