package main

import (
	"github.com/ReqqQ/eventpulse-go/src/infrastructure/server"
)

func main() {
	_, fiberApp := server.PrepareServerConfig()
	//controller.InitRoutes(fiberApp)
	server.StartServer(fiberApp)
}
