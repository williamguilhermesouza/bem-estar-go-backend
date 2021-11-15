package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/controllers/agenda"
)

func Agenda(app *fiber.App) {
	app.Get("/agenda", agenda.List)
	app.Get("/agenda/:id", agenda.FindOne)
	app.Post("/agenda", agenda.Create)
	app.Patch("/agenda/:id", agenda.Update)
	app.Delete("/agenda/:id", agenda.Delete)
}
