package attendances

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func List(c *fiber.Ctx) error {
	db := database.ConnectToDb()
	db.Table("attendance")
	attendance := []models.Attendance{}
	db.Find(&attendance)
	return c.JSON(attendance)
}
