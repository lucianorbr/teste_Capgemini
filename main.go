package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucianorbr/teste_Capgemini/controllers/functions"
	"github.com/lucianorbr/teste_Capgemini/db"
	"log"
)

func main() {

	app := fiber.New()

	app.Get("/stats", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"count_valid":   functions.CountIsValid(db.Letters),
			"count_invalid": functions.CountInValid(db.Letters),
			"ratio":         functions.CountRatio(db.Letters),
		})
	})

	app.Get("/sequence", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"letters": db.Letters,
		})
	})

	log.Fatal(app.Listen(":3000"))

}
