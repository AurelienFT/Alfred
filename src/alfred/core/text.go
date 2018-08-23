package core

import (

	"github.com/rthornton128/goncurses"

)

type Text struct {
	window *goncurses.Window
	panel *goncurses.Panel
	msg string
}

var texts = make([]*Text, 16)


func newText(msg string, x int, y int) {
	wind, _ := goncurses.NewWindow(1, len(msg), x, y)
	panel := goncurses.NewPanel(wind)
	wind.Printf(msg)
	texts = append(texts, &Text{window: wind, panel: panel, msg: msg})
}

func (class *Text) moveUp(x int) {
	class.panel.Move(0, x)
}

func moveUpText() {
	for text := range texts {
		texts[text].moveUp(1)
	}
}