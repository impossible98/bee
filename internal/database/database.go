package database

import (
	// import built-in packages
	"os"
	// import third-party packages
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// import local packages
	"bee/internal/model"
	"bee/pkg/path"
)

func InitDatabase() (*gorm.DB, error) {
	if !path.PathExists("data") {
		os.Mkdir("data", os.ModePerm)
	}
	database, err := gorm.Open(sqlite.Open("data/bee.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = database.AutoMigrate(&model.Favourite{})
	if err != nil {
		return nil, err
	}
	return database, nil
}
