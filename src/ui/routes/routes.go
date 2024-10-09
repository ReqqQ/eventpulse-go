package routes

import (
	"github.com/ReqqQ/eventpulse-user-go/src/app/user/dto"
	"github.com/ReqqQ/eventpulse-user-go/src/shared"
	_ "github.com/ThreeDotsLabs/watermill"
	_ "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	_ "github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofiber/fiber/v3"

	"github.com/ReqqQ/eventpulse-go/src/app/core/factory"
)

type Routes interface {
	InitRoutes(app *fiber.App)
}

type routesImpl struct {
	userApp     shared.Bus
	userFactory factory.UserFactory
}

func (r *routesImpl) InitRoutes(app *fiber.App) {
	app.Get("/user/dialog", func(c fiber.Ctx) error {
		loginDialogDTO := r.userApp.HandleQuery(r.userFactory.GetLoginDialogQuery(c.Query("socialType"))).GetData().(dto.GetLoginDialogDTO)

		return c.Redirect().To(loginDialogDTO.GetDialogUrl())
	})
	app.Get("/user/login", func(c fiber.Ctx) error {
		r.userApp.HandleCommand(r.userFactory.GetCreateUserFromSocialCommand(c.Query("code", "")))

		return c.Redirect().To("http://localhost:3000/")
	})

	//})

	//app.Get("/publisher", func(c fiber.Ctx) error {
	//	saramaPublisherConfig1 := kafka.DefaultSaramaSyncPublisherConfig()
	//	saramaPublisherConfig1.ClientID = "producer_1"
	//
	//	publisher1, err := kafka.NewPublisher(
	//		kafka.PublisherConfig{
	//			Brokers:               []string{"localhost:9092"},
	//			OverwriteSaramaConfig: saramaPublisherConfig1,
	//			Marshaler:             kafka.DefaultMarshaler{},
	//		},
	//		watermill.NewStdLogger(true, true),
	//	)
	//
	//	saramaPublisherConfig2 := kafka.DefaultSaramaSyncPublisherConfig()
	//	saramaPublisherConfig2.ClientID = "producer_2"
	//
	//	publisher2, err := kafka.NewPublisher(
	//		kafka.PublisherConfig{
	//			Brokers:               []string{"localhost:9092"},
	//			OverwriteSaramaConfig: saramaPublisherConfig2,
	//			Marshaler:             kafka.DefaultMarshaler{},
	//		},
	//		watermill.NewStdLogger(true, true),
	//	)
	//	if err != nil {
	//		return err
	//	}
	//	s := struct {
	//		UserId string
	//	}{
	//		UserId: "Patrykos",
	//	}
	//	jsonStruct, _ := json.Marshal(s)
	//
	//	err = publisher1.Publish("customUser", message.NewMessage(watermill.NewUUID(), jsonStruct))
	//	err = publisher2.Publish("customUser", message.NewMessage(watermill.NewUUID(), jsonStruct))
	//	if err != nil {
	//		return err
	//	}
	//
	//	return c.JSON("sended")
	//})

	//app.Get("/user", func(c fiber.Ctx) error {
	//	userApp := a.GetApps().GetUserApp()
	//	data := userApp.GetUserBus().HandleQuery(userApp.GetUserFactory().GetUserQuery(c.Query("userId")))
	//
	//	return c.JSON(data.GetData())
	//})

	//go cons()
}

func BuildRoutes(userApp shared.Bus, userFactory factory.UserFactory) Routes {
	return &routesImpl{
		userApp:     userApp,
		userFactory: userFactory,
	}
}

//func cons() {
//	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
//
//	subscriber, _ := kafka.NewSubscriber(
//		kafka.SubscriberConfig{
//			Brokers:               []string{"localhost:9092"},
//			Unmarshaler:           kafka.DefaultMarshaler{}, // JSON deserializator
//			ConsumerGroup:         "consumer_users",
//			OverwriteSaramaConfig: saramaSubscriberConfig,
//		},
//		watermill.NewStdLogger(false, false),
//	)
//	messagesUsers, _ := subscriber.Subscribe(context.Background(), "users")
//	for userData := range messagesUsers {
//
//		fmt.Printf("Offset: %d, Wiadomość: %s\n", "", string(userData.Payload))
//		userData.Ack()
//	}
//}
