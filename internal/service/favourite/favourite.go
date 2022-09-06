package favourite

import (
	// import local packages
	"bee/internal/model"
)

// Create Favourite
func CreateFavourite(identity string, platform string, roomId string, liveStreamer string) (*model.Favourite, error) {
	favourite := model.Favourite{
		Identity:     identity,
		Platform:     platform,
		RoomId:       roomId,
		LiveStreamer: liveStreamer,
		PlayURL:      platform + roomId,
	}
	if err := model.DB.Create(&favourite).Error; err != nil {
		return nil, err
	}
	return &favourite, nil
}
