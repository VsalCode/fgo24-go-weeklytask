package utils

import (
	"os"
	"os/exec"
)

func clear() {
	cmd := exec.Command("cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
