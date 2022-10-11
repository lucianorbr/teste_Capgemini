package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucianorbr/teste_Capgemini/controllers/functions"
	"github.com/lucianorbr/teste_Capgemini/db"
)

// Provide another endpoint "/stats" that responds to an HTTP GET. The answer must be a
// JSON that returns the stats of string checks
func Stats(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"count_valid":   functions.CountIsValid(db.Letters),
		"count_invalid": functions.CountInValid(db.Letters),
		"ratio":         functions.CountRatio(db.Letters),
	})
}
