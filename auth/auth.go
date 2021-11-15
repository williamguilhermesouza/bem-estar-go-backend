package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func Auth(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		// Users defines the allowed credentials
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
		// Authorizer defines a function you can pass
		// to check the credentials however you want.
		// It will be called with a username and password
		// and is expected to return true or false to indicate
		// that the credentials were approved or not.
		Authorizer: func(user, pass string) bool {
			if user == "john" && pass == "doe" {
				return true
			}
			if user == "admin" && pass == "123456" {
				return true
			}
			return false
		},
	}))
}
