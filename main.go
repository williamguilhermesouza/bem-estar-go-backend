package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamguilhermesouza/bem-estar-go-backend/auth"
	"github.com/williamguilhermesouza/bem-estar-go-backend/routes"
)

func main() {
	app := fiber.New()
	routes.Route(app)
	auth.Auth(app)
	app.Listen(":3003")
}
