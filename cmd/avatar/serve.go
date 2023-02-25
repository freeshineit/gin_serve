package avatar

import (
	"bytes"
	"fmt"
	"gin_serve/app/middleware"
	"gin_serve/config"
	"gin_serve/helper"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var AvatarCmd = &cobra.Command{
	Use:   "avatar",
	Short: "Run avatar serve",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		mode := cmd.Flag("mode").Value.String()
		Serve(mode)
	},
}

func Serve(mode string) error {
	r := gin.New()

	port := "9090"

	logger := config.SetUpZapLogger(mode == "release")

	// 中间件
	middleware.SetMiddleware(r, logger)

	r.GET("/avatar", BlueHandler)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	fmt.Printf("Avatar server listen: %s\n", helper.ColorBlueString("http://localhost:"+port))

	zap.S().Infof("Avatar server listen: %s\n", "http://localhost:"+port)

	err := helper.ListenAndServe(srv)

	if err != nil {
		zap.S().Fatal("Avatar server forced to shutdown:", err)
	}

	zap.S().Info("Avatar server exiting")

	defer logger.Sync()

	return err
}

func BlueHandler(ctx *gin.Context) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{0, 0, 255, 100}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.Point{
		X: 100,
		Y: 100,
	}, draw.Src)

	var img image.Image = m
	writeImage(ctx, &img)
}

// writeImage encodes an image 'img' in jpeg format and writes it into ResponseWriter.
func writeImage(ctx *gin.Context, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode image.")
	}

	ctx.Header("Content-Type", "image/png")
	// w.Header().Set("Content-Type", "application/octet-stream")
	ctx.Header("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := ctx.Writer.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
