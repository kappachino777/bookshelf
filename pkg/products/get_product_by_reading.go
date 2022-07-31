package products

import (
	"bookshelf/pkg/common/models"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetProductByReading(c *fiber.Ctx) error {
	reading := c.Params("reading")
	boolValue, _ := strconv.ParseBool(reading)

	var books []models.Book

	if result := h.DB.Where("reading = ?", boolValue).Find(&books); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&books)
}
