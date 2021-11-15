package rpgs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Update(c *fiber.Ctx) error {
	rpg := models.Rpg{}
	originalRpg := models.Rpg{}
	c.BodyParser(&rpg)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("rpg")
	db.First(&originalRpg, id)
	db.Model(&originalRpg).Updates(&rpg)
	return c.JSON(rpg)
}
