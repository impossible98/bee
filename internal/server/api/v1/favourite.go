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

// Get Favourite
func GetFavourite(ctx *gin.Context) {
	request := Favourite{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		format.HTTP(ctx, ecode.InvalidParams, err, nil)
		return
	}
	result, err := favourite.GetFavourite(request.Platform, request.RoomId)
	if err != nil {
		format.HTTP(ctx, ecode.ErrorGetFavourite, err, nil)
		return
	}
	format.HTTP(ctx, ecode.Success, nil, result)
}

// Get Favourite List
func GetFavouriteList(ctx *gin.Context) {
	result, err := favourite.GetFavouriteList()
	if err != nil {
		format.HTTP(ctx, ecode.ErrorGetFavouriteList, err, nil)
		return
	}
	format.HTTP(ctx, ecode.Success, nil, result)
}

// Delete Favourite
func DeleteFavourite(ctx *gin.Context) {
	request := Favourite{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		format.HTTP(ctx, ecode.InvalidParams, err, nil)
		return
	}
	result, err := favourite.DeleteFavourite(request.Platform, request.RoomId)
	if err != nil {
		format.HTTP(ctx, ecode.ErrorDeleteFavourite, err, nil)
		return
	}
	format.HTTP(ctx, ecode.Success, nil, result)
}
