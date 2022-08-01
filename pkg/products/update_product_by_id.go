package products

import (
	"time"

	"bookshelf/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateProductRequestBody struct {
	Name      string `json:"name"`
	Year      int    `json:"year"`
	Author    string `json:"author"`
	Summary   string `json:"summary"`
	Publisher string `json:"publisher"`
	PageCount int    `json:"pageCount"`
	ReadPage  int    `json:"readPage"`
	Reading   bool   `json:"reading"`
}

func (h handler) updateProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateProductRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if body.Name != "" {
		book.Name = body.Name
	}
	if body.Year != 0 {
		book.Year = body.Year
	}
	if body.Author != "" {
		book.Author = body.Author
	}
	if body.Summary != "" {
		book.Summary = body.Summary
	}
	if body.Publisher != "" {
		book.Publisher = body.Publisher
	}
	if body.PageCount != 0 {
		book.PageCount = body.PageCount
	}
	if body.ReadPage != 0 {
		book.ReadPage = body.ReadPage
	}

	book.Reading = body.Reading

	if book.PageCount == book.ReadPage {
		book.Finished = true
	} else {
		book.Finished = false
	}
	book.UpdatedAt = time.Now().Format("2017-09-07 17:06:06")

	h.DB.Save(&book)

	return c.Status(fiber.StatusCreated).JSON(&book)
}
