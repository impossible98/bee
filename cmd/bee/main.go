package main

import (
	// import local packages
	"bee/internal/database"
	"bee/internal/model"
	"bee/internal/server"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		panic(err)
	}
	model.DB = db
	server.InitServer()
}
