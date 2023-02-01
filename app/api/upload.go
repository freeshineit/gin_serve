package api

import (
	"gin_serve/app/utils"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// FileUpload file upload
// @Summary		Upload file
// @Description	Upload file
// @ID				file.upload
// @Accept			multipart/form-data
// @Produce		json
// @Param		file	formData	file true	"this is a test file"
// @Success		200		{object}	utils.BuildResponse
// @Failure		400		{object}	utils.BuildErrorResponse
// @Failure		404		{object}	utils.BuildErrorResponse
// @Router		/api/file_upload [post]
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

	c.JSON(http.StatusOK, utils.BuildResponse("success", gin.H{
		"urls": filesURL,
	}))
}
