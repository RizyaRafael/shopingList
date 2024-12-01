package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	product := app.Group("/products")
	ProductsRouter(product)

	user := app.Group("/user")
	UserRouter(user)
}
 