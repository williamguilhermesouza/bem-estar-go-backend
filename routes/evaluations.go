package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/controllers/evaluations"
)

func Evaluation(app *fiber.App) {
	app.Get("/evaluations", evaluations.List)
	app.Get("/evaluations/:id", evaluations.FindOne)
	app.Get("/evaluations/patientId/:id", evaluations.ByPatientId)
	app.Post("/evaluations", evaluations.Create)
	app.Patch("/evaluations/:id", evaluations.Update)
	app.Delete("/evaluations/:id", evaluations.Delete)
}
