package api

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"bee/global"
	"bee/pkg/ecode"
	"bee/pkg/format"
)

func Version(ctx *gin.Context) {
	format.HTTP(ctx, ecode.Success, nil, gin.H{
		"version": global.VERSION,
	})
}
