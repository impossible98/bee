package api

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"bee/pkg/ecode"
	"bee/pkg/format"
)

func API(ctx *gin.Context) {
	format.HTTP(ctx, ecode.Success, nil, gin.H{
		"version": "/api/version",
	})
}
