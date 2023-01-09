package api

import (
	"go_python_serve/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// FileUpload file upload
// /api/file_upload
func FileUpload(c *gin.Context) {

	filesURL := make([]string, 0)

	form, err := c.MultipartForm()

	log.Println(c.Cookie("token"))
	log.Println(c.GetHeader("Content-Type"))

	if err != nil {
		log.Printf("postMultipleFile error: %s \n", err.Error())
	}

	files := form.File["file"]

	_, err = os.Stat("upload")

	if err != nil {
		os.Mkdir("upload", os.ModePerm)
	}

	for _, file := range files {
		log.Println(file.Filename)

		// Upload the file to specific dst.
		if err = c.SaveUploadedFile(file, "upload/"+file.Filename); err != nil {
			log.Printf("SaveUploadedFile error: %s \n", err.Error())
			return
		}
		filesURL = append(filesURL, "images/upload/"+file.Filename)
	}

	c.JSON(http.StatusOK, models.BuildOKResponse(gin.H{
		"urls": filesURL,
	}))
}
