package init

import (
	c "github.com/ieee0824/thor/controller"
	"github.com/ieee0824/thor/view"
	termbox "github.com/nsf/termbox-go"
)

type ELBNameSetting struct {
	*TextBox
}

func (t *ELBNameSetting) Init() {
	t.TextBox = setELBName()
	t.SetController(setELBNameController)
}

func setELBName() *TextBox {
	return NewTextBox("使 用 す る ELB の 名 前 を 入 力 す る")
}

func setELBNameController() {
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
					if box.Answer() != "" {
						view.View["JP"].Fin()
						return
					}
				}
			case termbox.KeyBackspace:
				if box, err := view.View["JP"].GetView(); err == nil {
					if b, ok := box.(ELBNameSetting); ok {
						b.BS()
					}
					c.Draw(box)
				}
			case termbox.KeyBackspace2:
				if box, err := view.View["JP"].GetView(); err == nil {
					if b, ok := box.(ELBNameSetting); ok {
						b.BS()
					}
					c.Draw(box)
				}
			default:
				if box, err := view.View["JP"].GetView(); err == nil {
					if b, ok := box.(ELBNameSetting); ok {
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
