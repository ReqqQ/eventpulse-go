package ui

import (
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/gofiber/fiber/v3"
)

type FBTokenUser struct {
	AccessToken string `json:"access_token"`
}
type FBUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetUserToken struct {
	userId string
}

func (g GetUserToken) GetQueryUserId() string {
	return g.userId
}
func (g GetUserToken) GetType() string {
	return "GetUserToken"
}

var PubSub *gochannel.GoChannel

func (a *serverImpl) InitRoutes(app *fiber.App) {
	app.Get("/user/login", func(c fiber.Ctx) error {
		//link := a.GetUserController().GetAppService().LoginBySocial()

		return c.Redirect().To("")
	})
	app.Get("/user/get", func(c fiber.Ctx) error {
		//msg := message.NewMessage(watermill.NewUUID(), []byte(fmt.Sprintf("Wiadomość %d", i)))
		//err := PubSub.Publish("query.user", msg)
		//if err != nil {
		//	panic(err)
		//}
		//messages, err := PubSub.Subscribe("query.user.result")
		//if err != nil {
		//	return err
		//}
		//for msg := range messages {
		//	msg.Payload
		//}
		return c.JSON("get")
	})
	app.Get("/user/token", func(c fiber.Ctx) error {
		//client := resty.New()
		//fBTokenUser := new(FBTokenUser)
		//client.R().
		//	SetResult(&fBTokenUser).
		//	Get(a.GetUserController().GetAppService().GetToken(c.Query("code")))
		//fBUser := new(FBUser)
		//client.R().
		//	SetResult(&fBUser).
		//	SetHeader("Accept", "application/json").
		//	Get("https://graph.facebook.com/v20.0/me?fields=id,name&access_token=" + fBTokenUser.AccessToken)
		//
		//userUUID := uuid.NewV5(uuid.NamespaceURL, fBUser.Id)
		//userData := strings.Split(fBUser.Name, " ")
		//cluster := gocql.NewCluster("localhost")
		//cluster.Port = 9042
		//cluster.Keyspace = "omniloop"
		//cluster.Consistency = gocql.One
		//
		//session, err := cluster.CreateSession()
		//if err != nil {
		//	return err
		//}
		//uu, _ := gocql.ParseUUID(userUUID.String())
		//defer session.Close()
		//userEntity := models.UsersStruct{
		//	UserId:    uu,
		//	Name:      userData[0],
		//	Surname:   userData[1],
		//	CreatedAt: time.Now(),
		//}
		//stmt, names := qb.Insert("users").Columns("user_id", "name", "surname", "created_at").ToCql()
		//q := gocqlx.Query(session.Query(stmt), names).BindStruct(userEntity)
		//err = q.ExecRelease()
		//if err != nil {
		//	return err
		//}

		return c.Redirect().To("http://localhost:3000/")
	})
	app.Get("/user", func(c fiber.Ctx) error {
		data := a.GetApps().GetUserApp().GetUserBus().HandleQuery(GetUserToken{userId: c.Query("userId")})

		return c.JSON(data.GetData())
	})
	app.Post("/api/session", func(c fiber.Ctx) error {

		return c.JSON("Store")
	})
	app.Get("/api/session", func(c fiber.Ctx) error {
		//a.GetUserController().GetAppService().LoginBySocial()
		r := struct {
			Data string
		}{
			Data: "is logged",
		}
		return c.JSON(r)
	})
}
