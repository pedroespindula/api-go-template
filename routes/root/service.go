package root

import (
	"log"
	"os/exec"
)

func (r *Handler) service(nomeDoArquivo string) {
  log.Println("Running ROOT")

  cmd := exec.Command("/bin/sh", "-c", "touch output/" + nomeDoArquivo)

  err := cmd.Run()

  if err != nil {
    log.Fatal(err)
  }

  log.Println("Finished")
}
