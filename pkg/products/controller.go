package products

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/products")
	routes.Get("/:name", h.GetProductByName)
	routes.Get("/", h.GetProducts)
	routes.Get("/reading/:reading", h.GetProductByReading)
	routes.Get("/finished/:finished", h.GetProductByFinished)

	route := app.Group("/product")
	route.Delete("/:id", h.DeleteProductById)
	route.Get("/:id", h.GetProductById)
	route.Post("/", h.AddProduct)
	route.Put("/:id", h.updateProductById)
}
