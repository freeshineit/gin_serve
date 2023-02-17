package socket

import (
	"gin_serve/app/middleware"
	"gin_serve/app/routes"

	"gin_serve/config"
	"gin_serve/helper"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var SocketCmd = &cobra.Command{
	Use:   "socket",
	Short: "Run socket serve",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		mode := cmd.Flag("mode").Value.String()
		Serve(mode)
	},
}

func Serve(mode string) error {

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	port := config.Conf.Socket.Port

	r := gin.New()
	logger := config.SetUpZapLogger(mode == "release")

	middleware.SetMiddleware(r, logger)

	routes.SetupSocketRoutes(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Printf("socket server listen: %s\n", helper.ColorBlueString("http://localhost:"+port))

	err := helper.ListenAndServe(srv)

	if err != nil {
		zap.S().Fatalf("Socket Server forced to shutdown:", err)
	}

	log.Println("Socket Server exiting")

	return err
}
