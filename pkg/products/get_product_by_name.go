package products

import (
	"bookshelf/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetProductByName(c *fiber.Ctx) error {
	name := c.Params("name")

	var books []models.Book

	if result := h.DB.Where("name <> ?", name).Find(&books); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&books)
}
