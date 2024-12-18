package routes

import (
	"shopingList/controllers"
	"shopingList/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductsRouter(app fiber.Router) {
	app.Get("/", controllers.GetAllProducts)
	app.Get("/getOne/:id", controllers.GetOneProduct)

	//user needs to login to access
	app.Use(middleware.Authorization)
	app.Post("/create", controllers.CreateProduct)
	app.Post("/buyProduct", controllers.BuyProduct)
	app.Get("/getUserProducts", controllers.GetUserProducts)

	//Only user that created the product could access
	app.Put("/update/:id", middleware.Authentication, controllers.UpdateProduct)
	app.Delete("/delete/:id", middleware.Authentication, controllers.DeleteProduct)
}
