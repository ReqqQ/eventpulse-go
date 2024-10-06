package app

import (
	"github.com/ReqqQ/eventpulse-user-go/src/app/user"
	"github.com/ReqqQ/eventpulse-user-go/src/shared"
	"github.com/gofiber/fiber/v3"
)

type Server interface {
	InitRoutes(app *fiber.App)
	GetApps() App
}
type UserApp interface {
	GetUserBus() shared.Bus
	GetUserFactory() UserFactory
}

type appImpl struct {
	userApp UserApp
}

type UserFactory interface {
	GetUserQuery(userId string) user.GetUserQuery
}

type userAppImpl struct {
	bus     shared.Bus
	factory UserFactory
}

type App interface {
	GetUserApp() UserApp
}

func (u *userAppImpl) GetUserBus() shared.Bus {
	return u.bus
}
func (u *userAppImpl) GetUserFactory() UserFactory {
	return u.factory
}
func (a *appImpl) GetUserApp() UserApp {
	return a.userApp
}

func BuildApp(userApp UserApp) App {
	return &appImpl{userApp: userApp}
}

func BuildUserApp(bus shared.Bus, factory UserFactory) UserApp {
	return &userAppImpl{bus: bus, factory: factory}
}
