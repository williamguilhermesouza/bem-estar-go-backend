package movements

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Delete(c *fiber.Ctx) error {
	movement := models.Movement{}
	c.BodyParser(&movement)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("movement")
	db.First(&movement, id)
	db.Delete(&models.Movement{}, id)
	return c.JSON(movement)
}
