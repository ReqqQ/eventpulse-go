package main

import (
	"github.com/ReqqQ/eventpulse-go/src/infrastructure/server"
)

func main() {
	app := server.PrepareServerConfig()
	//controller := ui.GetControllersBase()
	//controller.InitRoutes(app)
	server.StartServer(app)
}
