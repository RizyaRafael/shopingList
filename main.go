package main

import (
	"log"
	"os"
	"shopingList/controllers"
	"shopingList/middleware"
	"shopingList/model"
	"shopingList/routes"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if os.Getenv("NODE_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

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
	middleware.DB = db

	app := fiber.New()
	app.Use(cors.New())
	routes.Routes(app)

	app.Listen(":3000")
}
