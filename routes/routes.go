package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/controllers/users"
)

func Route(app *fiber.App) {
	app.Get("/users", users.List)
	app.Post("/users", users.Create)
}
