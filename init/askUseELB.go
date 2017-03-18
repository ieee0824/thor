package init

import (
	c "github.com/ieee0824/thor/controller"
	"github.com/ieee0824/thor/util"
	"github.com/ieee0824/thor/view"
	termbox "github.com/nsf/termbox-go"
)

func init() {
}

type AskUseELB struct {
	*PolarQuestionBox
}

func (a *AskUseELB) Init() {
	a.PolarQuestionBox = askUseELBJP()
	a.SetController(askUseELBController)
}

func askUseELBJP() *PolarQuestionBox {
	return NewPolarQuestionBox(util.MultiString("ELBを使うかどうか"))
}

func askUseELBController() {
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
						// elbの設定へ
						view.View["JP"].Transition("elbNameSetting")
						return
					} else if box.Answer() == "no" {
						// ecsの設定へ
					}
				}
			case termbox.KeyBackspace:
				if box, err := view.View["JP"].GetView(); err == nil {
					if b, ok := box.(AskUseELB); ok {
						b.BS()
					}
					c.Draw(box)
				}
			case termbox.KeyBackspace2:
				if box, err := view.View["JP"].GetView(); err == nil {
					if b, ok := box.(AskUseELB); ok {
						b.BS()
					}
					c.Draw(box)
				}
			default:
				if box, err := view.View["JP"].GetView(); err == nil {
					if b, ok := box.(AskUseELB); ok {
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
