package core

/*type Text struct {
	window *goncurses.Window
	panel *goncurses.Panel
	msg string
	x_text int
}
var texts [100]*Text
var text_actual = 0

func NewText(msg string, x int, y int) {
	wind, _ := goncurses.NewWindow(1, len(msg), x, y)
	panel := goncurses.NewPanel(wind)
	wind.Printf(msg)
	texts[text_actual] = &Text{window: wind, panel: panel, msg: msg, x_text: x}
	text_actual++
}

func MoveUpText(x int) {
	var count = text_actual - 1
	for count >= 0 {
		texts[count].x_text -= x
		texts[count].panel.Move(texts[count].x_text, 0)
		count--
	}
}*/