package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetMiddleware(router *gin.Engine, logger *zap.Logger) {
	router.Use(GinZap(logger, time.RFC3339, true), RecoveryWithZap(logger, true), cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com", "http://localhost:8080", "http://localhost:8081", "http://localhost:8082", "http://localhost:3001", "http://localhost:3002", "http://localhost:3003", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
}
