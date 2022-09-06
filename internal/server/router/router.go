package router

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"bee/internal/server/api"
	"bee/internal/server/api/v1"
	"bee/internal/server/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())
	router.POST("/api", api.API)
	router.POST("/api/version", api.Version)
	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/favourite/create", v1.CreateFavourite)
		apiV1.POST("/favourite/get", v1.GetFavourite)
		apiV1.POST("/favourite_list/get", v1.GetFavouriteList)
	}
	return router
}
