package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/controllers/movements"
)

func Movements(app *fiber.App) {
	app.Get("/movements", movements.List)
	app.Get("/movements/:id", movements.FindOne)
	app.Post("/movements", movements.Create)
	app.Patch("/movements/:id", movements.Update)
	app.Delete("/movements/:id", movements.Delete)
}
