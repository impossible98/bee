package middleware

import (
	// import third-party packages
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST"},
		AllowHeaders:    []string{"Origin", "Content-Type"},
	})
}
