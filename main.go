package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucianorbr/teste_Capgemini/handlers"
	"log"
)

func main() {

	app := fiber.New()

	app.Get("/stats", handlers.Stats)
	app.Get("/sequence", handlers.Sequence)

	log.Fatal(app.Listen(":3000"))
}
