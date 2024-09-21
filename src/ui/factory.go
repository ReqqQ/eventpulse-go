package ui

import (
	"github.com/ReqqQ/eventpulse-user-go/src/app/user"
	"github.com/gofiber/fiber/v3"
)

type Controllers interface {
	GetUserController() UserController
	InitRoutes(app *fiber.App)
}
type UserController interface {
	GetAppService() user.UserService
}
type controllersImpl struct {
	userController UserController
}

func (c *controllersImpl) GetUserController() UserController {
	return c.userController
}

type userControllersImpl struct {
	userApp user.UserService
}

func (c *userControllersImpl) GetAppService() user.UserService {
	return c.userApp
}

func BuildControllers(uc UserController) Controllers {
	return &controllersImpl{userController: uc}
}

func BuildUserController(uc user.UserService) UserController {
	return &userControllersImpl{userApp: uc}
}
