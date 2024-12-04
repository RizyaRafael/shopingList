package routes

import (
	"shopingList/controllers"
	"shopingList/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductsRouter(app fiber.Router) {
	app.Get("/", controllers.GetAllProducts)

	//user needs to login to access
	app.Use(middleware.Authorization)
	app.Post("/create", controllers.CreateProduct)

	//Only user that created the product could access
	app.Use(middleware.Authentication)
	app.Put("/update", controllers.UpdateProduct)
	app.Delete("/delete", controllers.DeleteProduct)
}
