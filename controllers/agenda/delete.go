package agenda

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Delete(c *fiber.Ctx) error {
	agenda := models.Agenda{}
	c.BodyParser(&agenda)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("agenda")
	db.First(&agenda, id)
	db.Delete(&models.Agenda{}, id)
	return c.JSON(agenda)
}
