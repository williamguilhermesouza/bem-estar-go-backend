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
	originalUser.Name = user.Name
	originalUser.Sex = user.Sex
	originalUser.PhoneNumber = user.PhoneNumber
	originalUser.StreetName = user.StreetName
	originalUser.StreetNumber = user.StreetNumber
	originalUser.StreetDistrict = user.StreetDistrict
	originalUser.City = user.City
	originalUser.State = user.State
	originalUser.BirthDate = user.BirthDate
	originalUser.Cpf = user.Cpf
	originalUser.Email = user.Email
	db.Save(originalUser)
	return c.JSON(map[string]models.User{"user": user})
}
