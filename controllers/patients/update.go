package patients

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Update(c *fiber.Ctx) error {
	patient := models.Patient{}
	originalPatient := models.Patient{}
	c.BodyParser(&patient)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("patient")
	db.First(&originalPatient, id)
	db.Model(&originalPatient).Updates(&patient)
	return c.JSON(originalPatient)
}
