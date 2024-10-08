package routes

import (
	userDTO "github.com/ReqqQ/eventpulse-user-go/src/app/user/dto"
	"github.com/ReqqQ/eventpulse-user-go/src/shared"
	_ "github.com/ThreeDotsLabs/watermill"
	_ "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	_ "github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofiber/fiber/v3"

	"github.com/ReqqQ/eventpulse-go/src/app/core/factory"
)

type routes struct {
	userApp     shared.Bus
	userFactory factory.UserFactory
}

func (r *routes) InitRoutes(app *fiber.App) {
	app.Get("/user/dialog", func(c fiber.Ctx) error {
		dto := r.userApp.HandleQuery(r.userFactory.GetLoginDialogQuery(c.Query("socialType"))).(userDTO.GetLoginDialogDTO)

		return c.Redirect().To(dto.GetDialogUrl())
	})
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
	//app.Get("/user/login", func(c fiber.Ctx) error {
	//	code := c.Get("code", "")
	//	if code == "" {
	//		return c.Redirect().To("http://localhost:3000/login")
	//	}
	//
	//	userApp := a.GetApps().GetUserApp()
	//	userApp.GetUserBus()
	//	//_ := a.GetApps().GetUserApp()
	//	//userApp.GetUserBus().HandleCommand()
	//
	//	//client := resty.New()
	//	//fBTokenUser := new(FBTokenUser)
	//	//client.R().
	//	//	SetResult(&fBTokenUser).
	//	//	Get(a.GetUserController().GetAppService().GetToken(c.Query("code")))
	//	//fBUser := new(FBUser)
	//	//client.R().
	//	//	SetResult(&fBUser).
	//	//	SetHeader("Accept", "application/json").
	//	//	Get("https://graph.facebook.com/v20.0/me?fields=id,name&access_token=" + fBTokenUser.AccessToken)
	//	//
	//	//userUUID := uuid.NewV5(uuid.NamespaceURL, fBUser.Id)
	//	//userData := strings.Split(fBUser.Name, " ")
	//	//cluster := gocql.NewCluster("localhost")
	//	//cluster.Port = 9042
	//	//cluster.Keyspace = "omniloop"
	//	//cluster.Consistency = gocql.One
	//	//
	//	//session, err := cluster.CreateSession()
	//	//if err != nil {
	//	//	return err
	//	//}
	//	//uu, _ := gocql.ParseUUID(userUUID.String())
	//	//defer session.Close()
	//	//userEntity := models.UsersStruct{
	//	//	UserId:    uu,
	//	//	Name:      userData[0],
	//	//	Surname:   userData[1],
	//	//	CreatedAt: time.Now(),
	//	//}
	//	//stmt, names := qb.Insert("users").Columns("user_id", "name", "surname", "created_at").ToCql()
	//	//q := gocqlx.Query(session.Query(stmt), names).BindStruct(userEntity)
	//	//err = q.ExecRelease()
	//	//if err != nil {
	//	//	return err
	//	//}
	//
	//	return c.Redirect().To("http://localhost:3000/")
	//})
	//app.Get("/user", func(c fiber.Ctx) error {
	//	userApp := a.GetApps().GetUserApp()
	//	data := userApp.GetUserBus().HandleQuery(userApp.GetUserFactory().GetUserQuery(c.Query("userId")))
	//
	//	return c.JSON(data.GetData())
	//})

	//go cons()
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
