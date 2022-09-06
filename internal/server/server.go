package server

import (
	// import local packages
	"bee/internal/server/router"
)

func InitServer() {
	server := router.Router()
	server.Run(":7226")
}
