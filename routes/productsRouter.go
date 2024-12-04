package routes

import (
	"shopingList/controllers"
	"shopingList/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductsRouter(app fiber.Router) {
	app.Get("/", controllers.GetAllProducts)
	app.Post("/create", middleware.Authorization, controllers.CreateProduct)
	app.Put("/update", middleware.Authorization, middleware.Authentication, controllers.UpdateProduct)
}
