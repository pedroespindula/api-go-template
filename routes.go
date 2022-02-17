package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/pedroespindula/api-go-template/routes/end"
	"github.com/pedroespindula/api-go-template/routes/root"
)

type Handler interface {
  Method() string
  Route() string
  Handler() func(*fiber.Ctx) error
}

func routes() []Handler {
  return []Handler{
    end.NewHandler(),
    root.NewHandler(),
  }
}
