package app

import (
	"fmt"
	"gin_serve/app/middleware"
	"gin_serve/app/model"
	"gin_serve/app/routes"

	"gin_serve/config"
	"gin_serve/helper"
	"gin_serve/swagger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var AppCmd = &cobra.Command{
	Use:   "app",
	Short: "Run app serve",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		mode := cmd.Flag("mode").Value.String()
		Serve(mode)
	},
}

func Serve(mode string) error {

	// helper.InitTranslation("zh")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	port := config.Conf.App.Port

	logger := config.SetUpZapLogger(mode == "release")

	DB := config.SetupDatabaseConnection()
	config.SetupRedisConnection()

	model.GormAutoMigration(DB)

	r := gin.New()
	// 中间件
	middleware.SetMiddleware(r, logger)

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

	fmt.Printf("app server listen: %s\n", helper.ColorBlueString("http://localhost:"+port))

	zap.S().Infof("app server listen: %s\n", "http://localhost:"+port)

	err := helper.ListenAndServe(srv)

	if err != nil {
		zap.S().Fatal("App server forced to shutdown:", err)
	}

	zap.S().Info("App server exiting")

	defer logger.Sync()

	// close
	defer config.CloseMysqlConnection(DB)
	defer config.CloseRedisConnection(config.RedisClient)

	return err
}
