package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/controllers/patients"
	"github.com/williamguilhermesouza/bem-estar-go-backend/controllers/users"
)

func Route(app *fiber.App) {
	// users routes
	app.Get("/users", users.List)
	app.Post("/users", users.Create)
	// patients routes
	app.Get("/patients", patients.List)
}
