package init

import (
	c "github.com/ieee0824/thor/controller"
	"github.com/ieee0824/thor/util"
	"github.com/ieee0824/thor/view"
	termbox "github.com/nsf/termbox-go"
)

func init() {
	view.View["JP"] = view.New()
	generator := SelectGeneratorType{}
	generator.Init()
	view.View["JP"].Add(generator, "selecGeneratorTypeView")
	view.View["JP"].Transition("selecGeneratorTypeView")
}

type SelectGeneratorType struct {
	*SelectBox
}

func (s *SelectGeneratorType) Init() {
	s.SelectBox = selecGeneratorTypeJP()
	s.SetController(controller)
}

func controller() {
	for {
		if box, err := view.View["JP"].GetView(); err == nil {
			c.Draw(box)
		}
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyCtrlC:
				return
			case termbox.KeyArrowUp:
				if box, err := view.View["JP"].GetView(); err == nil {
					if b, ok := box.(SelectGeneratorType); ok {
						b.Up()
					}
					c.Draw(box)
				}
			case termbox.KeyArrowDown:
				if box, err := view.View["JP"].GetView(); err == nil {
					if b, ok := box.(SelectGeneratorType); ok {
						b.Down()
					}
					c.Draw(box)
				}

			default:
				if box, err := view.View["JP"].GetView(); err == nil {
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

func selecGeneratorTypeJP() *SelectBox {
	return NewSelectBox(
		util.MultiString("設定ファイルの生成方法を選ぶ"),
		[]string{
			util.MultiString("対話式に設定をする"),
			util.MultiString("テンプレートを生成する"),
		},
	)
}
