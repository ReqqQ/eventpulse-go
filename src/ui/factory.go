package ui

import (
	"github.com/ReqqQ/eventpulse-user-go/src/shared"
	"github.com/gofiber/fiber/v3"
)

type Server interface {
	InitRoutes(app *fiber.App)
	GetApps() App
}
type UserApp interface {
	GetUserBus() shared.Bus
}

type appImpl struct {
	userApp UserApp
}

type serverImpl struct {
	apps App
}
type userAppImpl struct {
	bus shared.Bus
}

type App interface {
	GetUserApp() UserApp
}

func (u *userAppImpl) GetUserBus() shared.Bus {
	return u.bus
}
func (a *appImpl) GetUserApp() UserApp {
	return a.userApp
}
func (s *serverImpl) GetApps() App {
	return s.apps
}
func BuildServer(apps App) Server {
	return &serverImpl{apps: apps}
}

func BuildApp(userApp UserApp) App {
	return &appImpl{userApp: userApp}
}

func BuildUserApp(bus shared.Bus) UserApp {
	return &userAppImpl{bus: bus}
}
