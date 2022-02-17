package end

import (
	"log"
	"os/exec"
)

func (h *Handler) service(nomeDoArquivo string) {
  log.Println("Running END")

  cmd := exec.Command("/bin/sh", "-c", "echo 'oi icaro' >> output/" + nomeDoArquivo)

  err := cmd.Run()

  if err != nil {
    log.Fatal(err)
  }

  log.Println("Finished")
}
