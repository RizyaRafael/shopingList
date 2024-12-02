package main

import (
	"log"
	"shopingList/controllers"
	"shopingList/model"
	"shopingList/routes"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=shoping_list port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.Products{})

	controllers.DB = db

	app := fiber.New()
	routes.Routes(app)

	app.Listen(":3000")
}
