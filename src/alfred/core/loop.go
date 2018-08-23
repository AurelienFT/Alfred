package core

import (
	//"bufio"
	//"os"
	"github.com/rthornton128/goncurses"
	"log"
)

var row int

// AlfredLoop is the main loop of Alfred
func AlfredLoop() {
	x := 3
	stdscr, err := goncurses.Init()
	if err != nil {
		log.Fatal("Cannot open ncurses window : ", err)
	}
	defer goncurses.End()
	welcome_msg := "Hello I'm Alfred your personnal assistant !"
	common_msg := "Can I help you ?"
	row, _ = stdscr.MaxYX()
	NewText(welcome_msg, row-3, 0)
	for x >= 0 {
		NewText(common_msg, row-2, 0)
		MoveUpText(row-x)
		goncurses.UpdatePanels()
		goncurses.Update()
		stdscr.GetChar()
		x++
	}
	goncurses.End()

}
