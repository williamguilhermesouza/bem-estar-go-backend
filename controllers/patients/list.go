package patients

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func List(c *fiber.Ctx) error {
	db := database.ConnectToDb()

	db.Table("patients")

	patient := []models.Patient{}

	db.Find(&patient)

	return c.JSON(map[string][]models.Patient{"patient": patient})
}
