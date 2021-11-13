package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func List(c *fiber.Ctx) error {
	db := database.ConnectToDb()

	db.Table("user")

	user := []models.User{}

	db.Find(&user)

	return c.JSON(map[string][]models.User{"user": user})
}
