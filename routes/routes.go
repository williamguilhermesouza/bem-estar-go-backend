package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	// users routes
	Users(app)
	// patients routes
	Patients(app)
}
