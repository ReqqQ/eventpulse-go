package main

import (
	"github.com/ReqqQ/eventpulse-go/src/infrastructure/server"
)

func main() {
	routes, fiberApp := server.PrepareServerConfig()
	routes.InitRoutes(fiberApp)
	server.StartServer(fiberApp)
}
