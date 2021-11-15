package attendances

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Update(c *fiber.Ctx) error {
	attendance := models.Attendance{}
	originalUser := models.Attendance{}
	c.BodyParser(&attendance)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("attendance")
	db.First(&originalUser, id)
	db.Model(&originalUser).Updates(&attendance)
	return c.JSON(attendance)
}
