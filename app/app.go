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
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	_ "gin_serve/app/utils"
)

var (
	// 连接mysql数据库
	DB          *gorm.DB
	RedisClient *redis.Client
)

func RunServer(conf config.ServerConfig) error {

	if conf.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	DB = config.SetupDatabaseConnection()
	RedisClient = config.SetupRedisConnection()

	model.GormAutoMigration(DB)

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
	defer config.CloseMysqlConnection(DB)
	defer config.CloseRedisConnection(RedisClient)

	return err
}
