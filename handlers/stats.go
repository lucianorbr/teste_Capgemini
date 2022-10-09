package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucianorbr/teste_Capgemini/controllers/functions"
	"github.com/lucianorbr/teste_Capgemini/db"
)

func Stats(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"count_valid":   functions.CountIsValid(db.Letters),
		"count_invalid": functions.CountInValid(db.Letters),
		"ratio":         functions.CountRatio(db.Letters),
	})
}