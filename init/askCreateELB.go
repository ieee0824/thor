package init

import (
	c "github.com/ieee0824/thor/controller"
	"github.com/ieee0824/thor/util"
	"github.com/ieee0824/thor/view"
	termbox "github.com/nsf/termbox-go"
)

func init() {
}

type AskCreateELB struct {
	*PolarQuestionBox
}

func (a *AskCreateELB) Init() {
	a.PolarQuestionBox = askCreateELBJP()
	a.SetController(askCreateELBController)
}

func askCreateELBJP() *PolarQuestionBox {
	return NewPolarQuestionBox(util.MultiString("ELBを作るかどうか"))
}

func askCreateELBController() {
	for {
		if box, err := view.View["JP"].GetView(); err == nil {
			c.Draw(box)
		}
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				view.View["JP"].Fin()
				return
			case termbox.KeyCtrlC:
				view.View["JP"].Fin()
				return
			case termbox.KeyEnter:
				if box, err := view.View["JP"].GetView(); err == nil {
					if box.Answer() == "y" {
					} else if box.Answer() == "no" {
					}
				}
			case termbox.KeyBackspace:
				if box, err := view.View["JP"].GetView(); err == nil {
					if b, ok := box.(AskCreateELB); ok {
						b.BS()
					}
					c.Draw(box)
				}
			default:
				if box, err := view.View["JP"].GetView(); err == nil {
					if b, ok := box.(AskCreateELB); ok {
						b.Add(ev.Ch)
					}
					c.Draw(box)
				}
			}
		default:
			if box, err := view.View["JP"].GetView(); err == nil {
				c.Draw(box)
			}
		}
	}
}
