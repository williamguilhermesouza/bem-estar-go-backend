package rpgs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Create(c *fiber.Ctx) error {
	rpg := models.Rpg{}
	c.BodyParser(&rpg)
	db := database.ConnectToDb()
	db.Table("rpg")
	db.Create(&rpg)
	return c.JSON(rpg)
}
