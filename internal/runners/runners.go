package runners

import (
	"log"
	"os/exec"
)

func Bash(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return "Failed to execute command", err
	}
	return string(output), nil
}
