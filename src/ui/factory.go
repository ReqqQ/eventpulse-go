package ui

import "github.com/ReqqQ/eventpulse-user-go/app/user"

type Controllers interface {
	GetUserController() UserController
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

//type Controllers interface {
//GetUserController() UserController
//InitRoutes(app *fiber.App)
//}
//type UserController interface {
//	CreateUser() struct {
//		make    string
//		model   string
//		mileage int
//	}
//}

//type controllersImpl struct {
//	userController UserController
//}
//
//type userControllerImpl struct {
//	app user.UserService
//}

//	func (u *userControllerImpl) GetUserApp() user.UserService {
//		return u.app
//	}
//
//	func (c *controllersImpl) GetUserController() UserController {
//		return c.userController
//	}
//func NewControllers(uc UserController) Controllers {
//	return &controllersImpl{userController: uc}
//}

//func NewUserController(app user.UserService) UserController {
//	return &userControllerImpl{app: app}
//}
