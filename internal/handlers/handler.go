package handlers

import (
	"main/internal/model"
	"main/internal/storage"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	storage *storage.Storage
}

func NewHandler(storage *storage.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}
func (h *Handler) GetNumber(c fiber.Ctx) error {
	nums := h.storage.GetArrayNum()
	return c.JSON(nums)
}

func (h *Handler) PostNum(c fiber.Ctx) error {
	num := new(model.Number)

	if err := c.Bind().Body(num); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.storage.PostNum(num.Num); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save number",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Number saved successfully",
	})
}
