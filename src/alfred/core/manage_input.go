package core

import (
	"fmt"
	"os/exec"
)

func manageInput(userInput string) {
	cmd := exec.Command("python", "src/alfred/neural_network/classify.py", userInput)
	out, _ := cmd.CombinedOutput()
	fmt.Print(string(out))
}
