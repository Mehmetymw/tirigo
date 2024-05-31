package handler

import (
	"tirigo/internal/dtos"
	"tirigo/internal/product-management/service"

	"github.com/gofiber/fiber"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var req dtos.ProductSaveParameter
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	product, err := h.service.Create(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}

func (h *ProductHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.service.Get(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(product)
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	products, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "product deleted"})
}

func (h *ProductHandler) InitializeRoutes(productRoutes fiber.Router) {
	productRoutes.Post("/", func(c *fiber.Ctx) {
		h.Create(c)
	})
	productRoutes.Get("/:id", func(c *fiber.Ctx) {
		h.Get(c)
	})
	productRoutes.Get("/", func(c *fiber.Ctx) {
		h.GetAll(c)
	})
	productRoutes.Delete("/:id", func(c *fiber.Ctx) {
		h.Delete(c)
	})
}
