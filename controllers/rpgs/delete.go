package rpgs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/database"
	"github.com/williamguilhermesouza/bem-estar-go-backend/models"
)

func Delete(c *fiber.Ctx) error {
	rpg := models.Rpg{}
	c.BodyParser(&rpg)
	id := c.Params("id")
	db := database.ConnectToDb()
	db.Table("rpg")
	db.First(&rpg, id)
	db.Delete(&models.Rpg{}, id)
	return c.JSON(rpg)
}
