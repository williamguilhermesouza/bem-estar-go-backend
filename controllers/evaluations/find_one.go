package evaluations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func FindOne(c *fiber.Ctx) error {
	evaluation := models.Evaluation{}
	c.BodyParser(&evaluation)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("evaluation")
	db.First(&evaluation, id)
	return c.JSON(evaluation)
}
