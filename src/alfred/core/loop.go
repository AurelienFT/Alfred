package core

import (
	//"bufio"
	//"os"
	"github.com/rthornton128/goncurses"
	"log"
)

// AlfredLoop is the main loop of Alfred
func AlfredLoop() {
	x := 3
	stdscr, err := goncurses.Init()
	if err != nil {
		log.Fatal("Cannot open ncurses window : ", err)
	}
	defer goncurses.End()
	msg := "Hello I'm Alfred your personnal assistant !"
	row, _ := stdscr.MaxYX()
	wind, _ := goncurses.NewWindow(1, len(msg), row-3, 0)
	panel := goncurses.NewPanel(wind)
	wind.Printf(msg)
	for x >= 0 {
		//need to add a tab with all wind and panel for do it for all messages and manage collision out of stdscr
		panel.Move(row-x, 0)
		goncurses.UpdatePanels()
		goncurses.Update()
		stdscr.GetChar()
		x++
	}
	goncurses.End()

}
