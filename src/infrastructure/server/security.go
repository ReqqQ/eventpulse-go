package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/encryptcookie"
	"github.com/gofiber/fiber/v3/middleware/session"
	"time"
)

var store = session.New(session.Config{
	CookieHTTPOnly: true,
	CookieSecure:   true,
	CookieSameSite: "Lax",
	Expiration:     24 * time.Hour,
})

func createSecurity(app *fiber.App) {
	corsInstance := cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost", "http://eventpulse.local"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	})
	cookieInstance := encryptcookie.New(encryptcookie.Config{
		Key: "9J2Tdvs12V0HVD10bCEVKzoGjrEEHRnpEGkftXmDbNY=",
	})

	app.Use(corsInstance, cookieInstance, apiMiddleware)
}
