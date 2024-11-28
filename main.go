package main

import (
	"log"
	"shopingList/model"
	"shopingList/routes"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=shoping_list port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.Products{})

	app := fiber.New()
	routes.Routes(app)

	app.Listen(":3000")
}
