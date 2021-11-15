package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	// users routes
	Users(app)
	// patients routes
	Patients(app)
	// attendances routes
	Attendances(app)
	// movements routes
	Movements(app)
	// agenda routes
	Agenda(app)
	// evaluations routes
	Evaluation(app)
	// rpgs routes
	Rpg(app)
}
