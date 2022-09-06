package v1

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// import local packages
	"bee/internal/service/favourite"
	"bee/pkg/ecode"
	"bee/pkg/format"
)

type Favourite struct {
	Identity     string `json:"identity"`
	Platform     string `json:"platform" binding:"required"`
	RoomId       string `json:"room_id" binding:"required"`
	LiveStreamer string `json:"live_streamer"`
}

// Create Favourite
func CreateFavourite(ctx *gin.Context) {
	request := Favourite{
		Identity: uuid.NewString(),
	}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		format.HTTP(ctx, ecode.InvalidParams, err, nil)
		return
	}
	result, err := favourite.CreateFavourite(request.Identity, request.Platform, request.RoomId, request.LiveStreamer)
	if err != nil {
		format.HTTP(ctx, ecode.ErrorCreateFavourite, err, nil)
		return
	}
	format.HTTP(ctx, ecode.Success, nil, result)
}
