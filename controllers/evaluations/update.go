package evaluations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Update(c *fiber.Ctx) error {
	evaluation := models.Evaluation{}
	originalEvaluation := models.Evaluation{}
	c.BodyParser(&evaluation)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("evaluation")
	db.First(&originalEvaluation, id)
	db.Model(&originalEvaluation).Updates(&evaluation)
	return c.JSON(evaluation)
}
