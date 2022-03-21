package main

import (
	"log"
	"news/database"
	"news/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)
	err := app.Listen("127.0.0.1:8001")
	if err != nil {
		log.Println(err)
	}
}
