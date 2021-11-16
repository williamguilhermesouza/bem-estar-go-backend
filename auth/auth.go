package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
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
			userModel := models.User{}
			db := database.ConnectToDb()
			db.Table("user")
			db.Where(map[string]interface{}{"email": user}).Find(&userModel)
			if user == userModel.Email && pass == userModel.Password {
				return true
			}
			return false
		},
	}))
}
