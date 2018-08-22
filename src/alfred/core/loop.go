package core

import (
	"bufio"
	"fmt"
	"os"
)

// AlfredLoop is the main loop of Alfred
func AlfredLoop() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hello I'm Alfred your peronal assistant !\n")
	for {
		fmt.Print("What can I do for you ?\n")
		text, _ := reader.ReadString('\n')
		fmt.Print(text)
	}
}
