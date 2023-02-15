package proxy

import (
	"gin_serve/app/middleware"
	"gin_serve/app/routes"
	"gin_serve/config"
	"gin_serve/helper"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var ProxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Run proxy serve",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		mode := cmd.Flag("mode").Value.String()
		Serve(mode)
	},
}

func Serve(mode string) error {

	r := gin.New()

	logger := config.SetUpZapLogger(mode == "release")

	middleware.SetMiddleware(r, logger)

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	port := config.Conf.Proxy.Port

	// Create a catchall route
	routes.SetupProxyRoutes(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Printf("proxy server listen: %s\n", helper.ColorBlueString("http://localhost:"+port))

	err := helper.ListenAndServe(srv)

	if err != nil {
		log.Fatal("Proxy server forced to shutdown:", err)
	}

	log.Println("Proxy server exiting")

	return err
}
