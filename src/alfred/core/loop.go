package core

import (
	"fmt"
	//"bufio"
	//"os"
	"github.com/rthornton128/goncurses"
	"log"
)

var row int
var col int
func B2S(bs []int) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}
// AlfredLoop is the main loop of Alfred
func AlfredLoop() {
	x := 1
	stdscr, err := goncurses.Init()
	if err != nil {
		log.Fatal("Cannot open ncurses window : ", err)
	}
	defer goncurses.End()
	welcome_msg := "Hello I'm Alfred your personnal assistant !"
	common_msg := "Can I help you ?"
	row, col = stdscr.MaxYX()
	stdscr.Move(row, 0)
	stdscr.Keypad(true)
	NewText(welcome_msg, row-1, 0)
	for x >= 0 {
		NewText(common_msg, row, 0)
		MoveUpText(2)
		goncurses.UpdatePanels()
		goncurses.Update()
		x++
		var key goncurses.Key
		key = goncurses.KEY_EOS
		stdscr.Move(row-1, 0)
		user_input := make([]int, 0)
		str := ""
		for key !=  10 {
			key = stdscr.GetChar()
			if key != 263 {
				user_input = append(user_input, int(key))
			} else {
				if (len(user_input) != 0) {
					user_input = user_input[:len(user_input)-1]
				}
				stdscr.DelChar()
			}
		}
		for _, i := range user_input {
			str = fmt.Sprintf("%s%c",str,i)
		}
	}
	goncurses.End()

}
