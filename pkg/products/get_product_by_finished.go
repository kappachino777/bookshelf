package products

import (
	"bookshelf/pkg/common/models"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetProductByFinished(c *fiber.Ctx) error {
	finished := c.Params("reafinishedding")
	boolValue, _ := strconv.ParseBool(finished)

	var books []models.Book

	if result := h.DB.Where("finished = ?", boolValue).Find(&books); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&books)
}
