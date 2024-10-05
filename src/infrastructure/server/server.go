package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"

	"github.com/ReqqQ/eventpulse-go/src/infrastructure/server/di"
	"github.com/ReqqQ/eventpulse-go/src/ui"
)

func PrepareServerConfig() (ui.Server, *fiber.App) {
	app := di.InitDIContainer()
	err := godotenv.Load(".env")
	if err != nil {
		return nil, nil
	}
	server := fiber.New()
	createSecurity(server)

	return app, server
}

func InitHandlers() {
	//router, err := message.NewRouter(message.RouterConfig{}, watermill.NewStdLogger(false, false))
	//if err != nil {
	//	panic(err)
	//}
	//ui.PubSub = gochannel.NewGoChannel(gochannel.Config{
	//	OutputChannelBuffer:            100,
	//	BlockPublishUntilSubscriberAck: false,
	//}, watermill.NewStdLogger(false, false))
	//
	//router.AddHandler(
	//	"userQueryHandler",
	//	"query.user",
	//	ui.PubSub,
	//	"query.user.result",
	//	ui.PubSub,
	//	new(user.QueryHandler).Handler,
	//)
	//
	//go func() {
	//	err := router.Run(context.Background())
	//	if err != nil {
	//		panic(err)
	//	}
	//}()
}

func StartServer(app *fiber.App) {
	app.Listen(":8080")
}
