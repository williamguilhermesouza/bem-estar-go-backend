package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/controllers/users"
)

func Users(app *fiber.App) {
	app.Get("/users", users.List)
	app.Get("/users/:id", users.FindOne)
	app.Post("/users", users.Create)
	app.Patch("/users/:id", users.Update)
	app.Delete("/users/:id", users.Delete)
}
