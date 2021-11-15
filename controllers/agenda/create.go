package agenda

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Create(c *fiber.Ctx) error {
	agenda := models.Agenda{}
	c.BodyParser(&agenda)
	db := database.ConnectToDb()
	db.Table("agenda")
	db.Create(&agenda)
	return c.JSON(agenda)
}
