package socket

import (
	"gin_serve/app/middleware"
	"gin_serve/app/routes"
	"gin_serve/config"
	"gin_serve/helper"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Serve(mode string) error {

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	port := config.Conf.Socket.Port

	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())

	routes.SetupSocketRoutes(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Printf("socket server listen: %s\n", helper.ColorBlueString("http://localhost:"+port))

	err := helper.ListenAndServe(srv)

	if err != nil {
		log.Fatal("Socket Server forced to shutdown:", err)
	}

	log.Println("Socket Server exiting")

	return err
}