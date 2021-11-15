package evaluations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Create(c *fiber.Ctx) error {
	evaluation := models.Evaluation{}
	c.BodyParser(&evaluation)
	db := database.ConnectToDb()
	db.Table("evaluation")
	db.Create(&evaluation)
	return c.JSON(evaluation)
}
