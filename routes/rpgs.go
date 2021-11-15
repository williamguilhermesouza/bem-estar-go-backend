package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/controllers/rpgs"
)

func Rpg(app *fiber.App) {
	app.Get("/rpgs", rpgs.List)
	app.Get("/rpgs/:id", rpgs.FindOne)
	app.Post("/rpgs", rpgs.Create)
	app.Patch("/rpgs/:id", rpgs.Update)
	app.Delete("/rpgs/:id", rpgs.Delete)
	app.Get("/rpgs/patientId/:id", rpgs.ByPatientId)

}
