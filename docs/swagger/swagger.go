package swagger

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitSwagger init swagger docs
func InitSwagger(r *gin.Engine) {

	r.Static("/docs", "./docs")

	url := ginSwagger.URL("http://localhost:8082/docs/swagger.json") // The url pointing to API definition

	log.Println("swagger api url \033[1;31;40mhttp://localhost:8082/swagger/index.html\033[0m")

	// http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
