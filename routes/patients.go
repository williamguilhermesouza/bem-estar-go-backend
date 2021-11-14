package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/controllers/patients"
)

func Patients(app *fiber.App) {
	// users routes
	app.Get("/patients", patients.List)
	app.Get("/patients/:id", patients.FindOne)
	app.Post("/patients", patients.Create)
	app.Patch("/patients/:id", patients.Update)
	app.Delete("/patients/:id", patients.Delete)
}
