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

// Get Favourite
func GetFavourite(platform string, roomId string) (*model.Favourite, error) {
	favourite := model.Favourite{
		PlayURL: platform + roomId,
	}
	if err := model.DB.Where("play_url = ?", favourite.PlayURL).First(&favourite).Error; err != nil {
		return nil, err
	}
	return &favourite, nil
}

// Get Favoutite List
func GetFavouriteList() ([]*model.Favourite, error) {
	favourites := *new([]*model.Favourite)
	if err := model.DB.Find(&favourites).Error; err != nil {
		return nil, err
	}
	return favourites, nil
}
