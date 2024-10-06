package ui

import (
	"github.com/ReqqQ/eventpulse-user-go/src/app/user"
	"github.com/gofiber/fiber/v3"
)

func (a *serverImpl) InitRoutes(app *fiber.App) {
	app.Get("/user/login", func(c fiber.Ctx) error {
		userApp := a.GetApps().GetUserApp()
		dto := userApp.GetUserBus().
			HandleQuery(
				userApp.GetUserFactory().GetLoginDialogQuery(c.Query("socialType"))).
			GetData().(user.GetLoginDialogDTO)

		return c.Redirect().To(dto.GetDialogUrl())
	})
	//app.Get("/user/token", func(c fiber.Ctx) error {
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
	app.Get("/user", func(c fiber.Ctx) error {
		userApp := a.GetApps().GetUserApp()
		data := userApp.GetUserBus().HandleQuery(userApp.GetUserFactory().GetUserQuery(c.Query("userId")))

		return c.JSON(data.GetData())
	})
}
