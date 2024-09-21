package server

import (
	"github.com/ReqqQ/eventpulse-go/src/infrastructure/server/di"
	"github.com/ReqqQ/eventpulse-go/src/ui"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func PrepareServerConfig() (ui.Controllers, *fiber.App) {
	app := di.InitDIContainer()
	err := godotenv.Load(".env")
	if err != nil {
		return nil, nil
	}
	server := fiber.New()
	//createSecurity(server)

	return app, server
}

func StartServer(app *fiber.App) {
	app.Listen(":8080")
}
