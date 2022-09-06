package model

import (
	// import third-party packages
	"gorm.io/gorm"
)

type Favourite struct {
	gorm.Model
	Identity     string `gorm:"column:identity;not null" json:"identity"`
	Platform     string `gorm:"column:platform;not null" json:"platform"`
	RoomId       string `gorm:"column:room_id;not null" json:"room_id"`
	LiveStreamer string `gorm:"column:live_streamer;not null" json:"live_streamer"`
	PlayURL      string `gorm:"column:play_url;not null;unique" json:"play_url"`
}
