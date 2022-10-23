package main

import (
	"log"

	"github.com/Paul-Boersma/go-div-rhino/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
