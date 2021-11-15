package evaluations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func List(c *fiber.Ctx) error {
	db := database.ConnectToDb()
	db.Table("evaluation")
	evaluation := []models.Evaluation{}
	db.Find(&evaluation)
	return c.JSON(evaluation)
}
