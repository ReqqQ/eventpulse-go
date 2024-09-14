package ui

//func (cb *controllersImpl) InitRoutes(app *fiber.App) {
//	//app.Get("api/", func(c fiber.Ctx) error {
//	//	return c.JSON(cb.GetUserController().CreateUser())
//	//})
//	//app.Get("api/login/:name", func(c fiber.Ctx) error {
//	//	userService := &UserApp.UserServiceImpl{}
//	//	userService.GetUser()
//	//
//	//	param := c.Params("name", "")
//	//	if param == "" {
//	//		return c.SendString("No name")
//	//	}
//	//	sess, err := store.Get(c)
//	//	if err != nil {
//	//		return err
//	//	}
//	//	state := sess.Get("userState")
//	//
//	//	if state != nil && state != param {
//	//		return c.SendString("Wrong user. Skipping")
//	//	}
//	//
//	//	if state != nil {
//	//		return c.SendString("Logged")
//	//	}
//	//	sess.SetExpiry(10 * time.Second)
//	//	sess.Set("userState", param)
//	//	sess.Save()
//	//	return c.SendString("Logged")
//	//})
//}
