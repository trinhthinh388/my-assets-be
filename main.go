package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"my-assets-be/routes"
)

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("Error while loading .env file.")
	}

	app := fiber.New()

	v1 := app.Group("/v1")

	routes.AccountRoutes(v1)

	app.Listen(":3000")
}
