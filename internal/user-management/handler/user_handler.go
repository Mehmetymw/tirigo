package handler

import (
	"tirigo/internal/user-management/dtos"
	"tirigo/internal/user-management/service"

	"github.com/gofiber/fiber"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {

	var req dtos.UserRegisterParameter
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.service.Register(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) Authenticate(c *fiber.Ctx) error {
	var req dtos.UserAuthParameter
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.service.Authenticate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func (h *UserHandler) Logout(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing token"})
	}
	err := h.service.Logout(token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "logged out"})
}

func (h *UserHandler) InitializeRoutes(userRoutes fiber.Router) {

	userRoutes.Post("/register", func(c *fiber.Ctx) {
		h.Register(c)
	})
	userRoutes.Post("/login", func(c *fiber.Ctx) {
		h.Authenticate(c)
	})

	userRoutes.Post("/logout", func(c *fiber.Ctx) {
		h.Logout(c)
	})

}
