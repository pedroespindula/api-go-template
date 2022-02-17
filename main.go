package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
  api := fiber.New()

  for _, r := range routes() {
    api.Add(r.Method(), r.Route(), r.Handler())
  }

  api.Listen(":3000")
}
