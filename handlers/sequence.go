package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucianorbr/teste_Capgemini/db"
)

// Returns an endpoint "/sequence" in JSON format with all strings of Letters stored in the database
func Sequence(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"letters": db.Letters,
	})
}
