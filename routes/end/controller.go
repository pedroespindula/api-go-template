package end

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Body struct {
  NomeDoArquivo string `json:"nome_do_arquivo"`
}

func (h *Handler) controller(c *fiber.Ctx) error {
  body := new(Body)  
  if err := c.BodyParser(body); err != nil {
    return errors.New("Deu merda")
  }

  h.service(body.NomeDoArquivo)

  return nil
}

