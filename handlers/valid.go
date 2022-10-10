package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucianorbr/teste_Capgemini/controllers/functions"
	"github.com/lucianorbr/teste_Capgemini/db"
)

func Valid(c *fiber.Ctx) error {
	for i := 0; i < len(db.Letters); i++ {
		//fmt.Printf("{\"is_valid\": %v}\n", functions.IsValid(db.Letters[i]))

		//return all JSON Letters

		//return c.JSON(fmt.Sprintf("{\"is_valid\": %v}", functions.IsValid(db.Letters[i])))

		return c.JSON(fiber.Map{
			"is_valid": functions.IsValid(db.Letters[i]),
		})
	}

	return nil
}
