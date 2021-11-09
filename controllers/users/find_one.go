package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func FindOne(c *fiber.Ctx) error {
	user := models.User{}
	c.BodyParser(&user)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("user")
	db.First(&user, id)
	return c.JSON(map[string]models.User{"user": user})
}
