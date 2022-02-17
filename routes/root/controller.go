package root

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Body struct {
  NomeDoArquivo string `json:"nome_do_arquivo"`
}

func (r *Handler) controller(c *fiber.Ctx) error {
  body := new(Body)  
  if err := c.BodyParser(body); err != nil {
    return errors.New("An error has occurred")
  }

  r.service(body.NomeDoArquivo)

  return nil
}

