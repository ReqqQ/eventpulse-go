package server

import (
	"github.com/ReqqQ/eventpulse-user-go/src/shared"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"

	"github.com/ReqqQ/eventpulse-go/config/di"
	"github.com/ReqqQ/eventpulse-go/src/app"
)

func PrepareServerConfig() (app.Server, *fiber.App) {
	container := di.InitDIContainer()
	prepareDB()
	err := godotenv.Load(".env")
	if err != nil {
		return nil, nil
	}
	server := fiber.New()
	createSecurity(server)

	return container, server
}

func prepareDB() {
	cluster := gocql.NewCluster("localhost")
	cluster.Port = 9042
	cluster.Keyspace = "omniloop"
	cluster.Consistency = gocql.One

	shared.SetDBCluster(cluster)
}

func StartServer(app *fiber.App) {
	app.Listen(":8080")
}
