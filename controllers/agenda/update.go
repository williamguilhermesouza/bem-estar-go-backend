package agenda

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Update(c *fiber.Ctx) error {
	agenda := models.Agenda{}
	originalAgenda := models.Agenda{}
	c.BodyParser(&agenda)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("agenda")
	db.First(&originalAgenda, id)
	db.Model(&originalAgenda).Updates(&agenda)
	return c.JSON(agenda)
}
