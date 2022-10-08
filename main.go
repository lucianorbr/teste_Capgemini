package main

import (
	"fmt"
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

//var resultisvalid = functions.CountIsValid(letters)
//var resultratio = functions.CountRatio(letters)
//var resultinvalid = functions.CountInValid(letters)

//fmt.Printf("{\"is_valid\": %v}\n", functions.IsValid(letters[i]))

func main() {

	for i := 0; i < len(letters); i++ {
		fmt.Printf("{\"is_valid\": %v}\n", functions.IsValid(letters[i]))
	}

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
