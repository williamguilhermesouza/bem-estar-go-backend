package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/controllers/attendances"
)

func Attendances(app *fiber.App) {
	app.Get("/attendances", attendances.List)
	app.Get("/attendances/:id", attendances.FindOne)
	app.Get("/attendances/patientId/:id", attendances.ByPatientId)
	app.Post("/attendances", attendances.Create)
	app.Patch("/attendances/:id", attendances.Update)
	app.Delete("/attendances/:id", attendances.Delete)
}
