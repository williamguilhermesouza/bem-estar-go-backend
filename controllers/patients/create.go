package patients

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Create(c *fiber.Ctx) error {
	patient := models.Patient{}

	c.BodyParser(&patient)

	db := database.ConnectToDb()

	db.Table("patient")

	db.Create(&patient)

	return c.JSON(map[string]models.Patient{"patient": patient})
}
