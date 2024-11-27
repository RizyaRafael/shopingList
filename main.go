package main

import (
	"shopingList/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Routes(app)

	app.Listen(":3000")
}
