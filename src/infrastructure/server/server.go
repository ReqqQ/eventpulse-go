package server

import (
	di "github.com/ReqqQ/eventpulse-go/src/infrastructure/server/di"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func PrepareServerConfig() *fiber.App {
	app := di.InitDIContainer()

	app.GetUserController().GetAppService().GetUser()
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}
	server := fiber.New()
	//createSecurity(server)

	return server
}

func StartServer(app *fiber.App) {
	app.Listen(":8080")
}
