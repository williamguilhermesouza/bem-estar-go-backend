package evaluations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func ByPatientId(c *fiber.Ctx) error {
	evaluation := models.Evaluation{}
	c.BodyParser(&evaluation)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("evaluation")
	db.Where(map[string]interface{}{"patientId": id}).Find(&evaluation)
	return c.JSON(evaluation)
}
