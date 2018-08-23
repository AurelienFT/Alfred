package core

import (

	"github.com/rthornton128/goncurses"
	"fmt"

)

type Text struct {
	window *goncurses.Window
	panel *goncurses.Panel
	msg string
}
var texts [21]*Text
var text_actual = 0

func get_row_value() int {
	return row
}

func NewText(msg string, x int, y int) {
	fmt.Printf("%d\n", row)
	wind, _ := goncurses.NewWindow(1, len(msg), x, y)
	panel := goncurses.NewPanel(wind)
	wind.Printf(msg)
	texts[text_actual] = &Text{window: wind, panel: panel, msg: msg}
	text_actual++
}

func MoveUpText(x int) {
	for text := range texts {
		texts[text].panel.Move(x, 0)
	}
}