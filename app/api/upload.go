package api

import (
	"gin_serve/helper"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// FileUpload file upload
// @Summary		Upload file
// @Description	Upload file
// @Tags	    Upload
// @ID				file.upload
// @Accept			multipart/form-data
// @Produce		json
// @Param		file	formData	file true	"upload image"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Router		/api/file_upload [post]
// @Security    Bearer
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

	c.JSON(http.StatusOK, helper.BuildResponse("success", gin.H{
		"urls": filesURL,
	}))
}
