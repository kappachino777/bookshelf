package products

import (
	"bookshelf/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&book)
}
