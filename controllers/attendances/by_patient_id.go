package attendances

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func ByPatientId(c *fiber.Ctx) error {
	attendance := models.Attendance{}
	c.BodyParser(&attendance)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("attendance")
	db.Where(map[string]interface{}{"patientId": id}).Find(&attendance)
	return c.JSON(attendance)
}
