package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucianorbr/teste_Capgemini/db"
)

func Sequence(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"letters": db.Letters,
	})
}
