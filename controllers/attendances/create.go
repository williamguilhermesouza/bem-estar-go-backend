package attendances

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Create(c *fiber.Ctx) error {
	attendance := models.Attendance{}
	c.BodyParser(&attendance)
	db := database.ConnectToDb()
	db.Table("attendance")
	db.Create(&attendance)
	return c.JSON(attendance)
}
