package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucianorbr/teste_Capgemini/controllers/functions"
	"github.com/lucianorbr/teste_Capgemini/db"
)

func Valid(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{

		"DUHBHB - is_valid": functions.IsValid(db.Letters[0]),
		"DUBUHD - is_valid": functions.IsValid(db.Letters[1]),
		"UBUUHU - is_valid": functions.IsValid(db.Letters[2]),
		"BHBDHH - is_valid": functions.IsValid(db.Letters[3]),
		"DDDDUB - is_valid": functions.IsValid(db.Letters[4]),
		"UDBDUH - is_valid": functions.IsValid(db.Letters[5]),
		"OOOOOO - is_valid": functions.IsValid(db.Letters[6]),
		"MMMMMN - is_valid": functions.IsValid(db.Letters[7]),
		"NNNNNN - is_valid": functions.IsValid(db.Letters[8]),
		"HHHHHH - is_valid": functions.IsValid(db.Letters[9]),
		"YYYYYY - is_valid": functions.IsValid(db.Letters[10]),
		"XXXXXX - is_valid": functions.IsValid(db.Letters[11]),
	})
}
