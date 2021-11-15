package movements

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Create(c *fiber.Ctx) error {
	movement := models.Movement{}
	c.BodyParser(&movement)
	db := database.ConnectToDb()
	db.Table("movement")
	db.Create(&movement)
	return c.JSON(movement)
}
