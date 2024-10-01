package ui

import (
	"context"
	"database/sql"
	models "github.com/ReqqQ/eventpulse-user-go/src/infrastructure/users/ORM"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v3"
	"github.com/gofrs/uuid"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"strings"
	"time"
)

type FBTokenUser struct {
	AccessToken string `json:"access_token"`
}
type FBUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (cb *controllersImpl) InitRoutes(app *fiber.App) {
	app.Get("/user/login", func(c fiber.Ctx) error {
		link := cb.GetUserController().GetAppService().LoginBySocial()

		return c.Redirect().To(link)
	})
	app.Get("/user/token", func(c fiber.Ctx) error {
		client := resty.New()
		client2 := client.Clone()
		//t := c.Queries()
		//fmt.Println(t)
		fBTokenUser := new(FBTokenUser)
		client.R().
			SetResult(&fBTokenUser).
			Get(cb.GetUserController().GetAppService().GetToken(c.Query("code")))
		fBUser := new(FBUser)
		client2.R().
			SetResult(&fBUser).
			SetHeader("Accept", "application/json").
			Get("https://graph.facebook.com/v20.0/me?fields=id,name&access_token=" + fBTokenUser.AccessToken)

		db, err := sql.Open("postgres", "postgres://eventpulse:eventpulsepassword@localhost:5432/eventpulse?sslmode=disable")
		if err != nil {
			return err
		}
		boil.SetDB(db)
		userUUID := uuid.NewV5(uuid.NamespaceURL, fBUser.Id)
		if err != nil {
			return err
		}
		userExists, err := models.Users(models.UserWhere.ID.EQ(userUUID.String())).ExistsG(context.Background())
		if err != nil {
			return err
		}
		if userExists {
			return c.Redirect().To("http://localhost:3000/")
		}
		userData := strings.Split(fBUser.Name, " ")

		user := models.User{
			ID:        userUUID.String(),
			Name:      userData[0],
			Surname:   null.StringFrom(userData[1]),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		user.InsertG(context.Background(), boil.Infer())
		return c.Redirect().To("http://localhost:3000/")
	})
	app.Get("/api/user", func(c fiber.Ctx) error {
		db, err := sql.Open("postgres", "postgres://eventpulse:eventpulsepassword@localhost:5432/eventpulse?sslmode=disable")
		if err != nil {
			return err
		}

		boil.SetDB(db)

		users, err := models.Users().AllG(context.Background())

		return c.JSON(users)
	})
	app.Post("/api/session", func(c fiber.Ctx) error {

		return c.JSON("Store")
	})
	app.Get("/api/session", func(c fiber.Ctx) error {
		//cb.GetUserController().GetAppService().LoginBySocial()
		r := struct {
			Data string
		}{
			Data: "is logged",
		}
		return c.JSON(r)
	})
}
