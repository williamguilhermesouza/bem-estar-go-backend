package patients

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func FindOne(c *fiber.Ctx) error {
	patient := models.Patient{}
	c.BodyParser(&patient)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("patient")
	db.First(&patient, id)
	return c.JSON(map[string]models.Patient{"patient": patient})
}
