package routes

import (
	"shopingList/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}
