package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Create(c *fiber.Ctx) error {
	user := models.User{}
	c.BodyParser(&user)
	db := database.ConnectToDb()
	db.Table("user")
	db.Create(&user)
	return c.JSON(user)
}
