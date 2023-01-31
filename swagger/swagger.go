package swagger

import (
	"gin_serve/helpers"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// InitSwagger init swagger docs
func InitSwagger(r *gin.Engine) {

	r.Static("/docs", "./docs")

	url := ginSwagger.URL("http://localhost:8080/docs/swagger.json") // The url pointing to API definition

	log.Printf("swagger docs api url %s \n", helpers.ColorBlueString("http://localhost:8080/swagger/index.html"))
	// http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
