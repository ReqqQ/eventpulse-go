package main

import (
	"github.com/ReqqQ/eventpulse-go/src/infrastructure/server"
)

func main() {
	controller, fiberApp := server.PrepareServerConfig()
	server.InitHandlers()
	controller.InitRoutes(fiberApp)
	server.StartServer(fiberApp)
}
