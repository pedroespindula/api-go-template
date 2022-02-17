package root

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {}

func NewHandler() *Handler {
  return &Handler{}
}

func (h *Handler) Route() string {
  return "/"
}

func (h *Handler) Method() string {
  return "POST"
}

func (h *Handler) Handler() func(*fiber.Ctx) error {
  return h.controller
}
