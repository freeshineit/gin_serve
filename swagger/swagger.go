package swagger

import (
	"gin_serve/helper"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwagger(r *gin.Engine) {

	r.Static("/docs", "./docs")

	url := ginSwagger.URL("http://localhost:8080/docs/swagger.json") // The url pointing to API definition

	log.Printf("swagger docs api url %s \n", helper.ColorBlueString("http://localhost:8080/swagger/index.html"))
	// http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
