package movements

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Update(c *fiber.Ctx) error {
	movement := models.Movement{}
	originalMovement := models.Movement{}
	c.BodyParser(&movement)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("movement")
	db.First(&originalMovement, id)
	db.Model(&originalMovement).Updates(&movement)
	return c.JSON(movement)
}
