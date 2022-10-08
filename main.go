package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucianorbr/teste_Capgemini/controllers/functions"
	"log"
)

var letters = []string{
	"DUHBHB",
	"DUBUHD",
	"UBUUHU",
	"BHBDHH",
	"DDDDUB",
	"UDBDUH",
	"NNNNNN",
	"MMMMMN",
	"NNNNNN",
	"HHHHHH",
	"YYYYYY",
	"XXXXXX",
}

func main() {

	app := fiber.New()

	app.Get("/stats", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"count_valid":   functions.CountIsValid(letters),
			"count_invalid": functions.CountInValid(letters),
			"ratio":         functions.CountRatio(letters),
		})
	})

	app.Get("/sequence", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"letters": letters,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
