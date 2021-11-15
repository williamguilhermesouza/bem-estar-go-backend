package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Update(c *fiber.Ctx) error {
	user := models.User{}
	originalUser := models.User{}
	c.BodyParser(&user)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("user")
	db.First(&originalUser, id)
	db.Model(&originalUser).Updates(&user)
	return c.JSON(user)
}
