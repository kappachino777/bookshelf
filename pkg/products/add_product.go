package products

import (
	"time"

	"bookshelf/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type AddProductRequestBody struct {
	Name      string `json:"name"`
	Year      int    `json:"year"`
	Author    string `json:"author"`
	Summary   string `json:"summary"`
	Publisher string `json:"publisher"`
	PageCount int    `json:"pageCount"`
	ReadPage  int    `json:"readPage"`
	Reading   bool   `json:"reading"`
}

func (h handler) AddProduct(c *fiber.Ctx) error {
	body := AddProductRequestBody{}

	// parse body, attach to AddProductRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	book.Name = body.Name
	book.Year = body.Year
	book.Author = body.Author
	book.Summary = body.Summary
	book.Publisher = body.Publisher
	book.PageCount = body.PageCount
	book.ReadPage = body.ReadPage
	book.Reading = body.Reading
	if book.PageCount == book.ReadPage {
		book.Finished = true
	} else {
		book.Finished = false
	}
	book.InsertedAt = time.Now().Format("2017-09-07 17:06:06")
	book.UpdatedAt = book.InsertedAt

	// insert new db entry
	if result := h.DB.Create(&book); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&book)
}
