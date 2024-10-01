package server

import (
	"github.com/gofiber/fiber/v3"
)

func apiMiddleware(c fiber.Ctx) error {
	//println("Middleware: Żądanie otrzymane")
	//
	//if strings.Contains(c.OriginalURL(), "/api/login") || strings.Contains(c.OriginalURL(), "/favicon") {
	//	return c.Next()
	//}
	//
	//sess, _ := store.Get(c)
	//state := sess.Get("userState")
	//
	//if state == nil {
	//	c.Response().SetStatusCode(400)
	//
	//	response := struct {
	//		Message string `json:"message"`
	//	}{Message: "No logged user"}
	//	return c.JSON(response)
	//}
	//println("Hello, World! Witaj: " + fmt.Sprint(sess.Get("userState")))

	return c.Next()
}
