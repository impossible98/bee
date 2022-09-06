package router

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"bee/internal/server/api"
	"bee/internal/server/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())
	router.POST("/api/version", api.Version)
	return router
}
