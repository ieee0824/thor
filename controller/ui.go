package controller

import (
	"strings"

	"github.com/ieee0824/thor/view"
	termbox "github.com/nsf/termbox-go"
)

func drawLine(x, y int, str string) {
	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)
	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func DrawString(s string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawLine(0, 0, "Press ^C or ESC to exit.")
	for i, s := range strings.Split(s, "\n") {
		drawLine(0, 1+i, s)
	}
	termbox.Flush()
}

func Draw(d view.SingleView) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawLine(0, 0, "Press ^C or ESC to exit.")

	for i, l := range strings.Split(d.Message(), "\n") {
		drawLine(0, i+1, l)
	}

	termbox.Flush()
}
