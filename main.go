package main

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/encryptcookie"
	"github.com/gofiber/fiber/v3/middleware/session"
	"strings"
	"time"
)

var store = session.New(session.Config{
	CookieHTTPOnly: true,           // Ciasteczko HttpOnly - niedostępne dla JS
	CookieSecure:   true,           // Ciasteczko dostępne tylko przez HTTPS
	CookieSameSite: "Lax",          // Ochrona przed atakami CSRF
	Expiration:     24 * time.Hour, // Czas ważności sesji
})

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost", "http://eventpulse.local"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))
	app.Use(
		encryptcookie.New(encryptcookie.Config{
			Key: "",
		}),
		myMiddleware)
	app.Get("api/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("api/login/:name", func(c fiber.Ctx) error {
		param := c.Params("name", "")
		if param == "" {
			return c.SendString("No name")
		}
		sess, err := store.Get(c)
		if err != nil {
			return err
		}
		state := sess.Get("userState")

		if state != nil && state != param {
			return c.SendString("Wrong user. Skipping")
		}

		if state != nil {
			return c.SendString("Logged")
		}
		sess.Set("userState", param)
		sess.Save()
		return c.SendString("Logged")
	})
	app.Listen(":8080")
}

func myMiddleware(c fiber.Ctx) error {
	println("Middleware: Żądanie otrzymane")

	if strings.Contains(c.OriginalURL(), "/api/login") {
		return c.Next()
	}

	sess, _ := store.Get(c)
	state := sess.Get("userState")

	if state == nil {
		return c.SendString("No logged user")
	}
	println("Hello, World! Witaj: " + fmt.Sprint(sess.Get("userState")))

	return c.Next()
}
