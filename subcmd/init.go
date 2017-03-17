package subcmd

import (
	"log"
	"time"

	"github.com/ieee0824/thor/controller"
	_ "github.com/ieee0824/thor/init"
	"github.com/ieee0824/thor/view"
	termbox "github.com/nsf/termbox-go"
)

func init() {
}

type Init struct{}

func (c *Init) Run(args []string) int {
	err := termbox.Init()
	if err != nil {
		log.Fatalln(err)
	}
	defer termbox.Close()
	go func() {
		t := time.NewTicker(500 * time.Millisecond)
		for {
			select {
			case <-t.C:
				if box, err := view.View["JP"].GetView(); err == nil {
					box.ToggleCursor()
					controller.Draw(box)
				}
			default:
			}
		}
		t.Stop()
	}()
	if v, err := view.View["JP"].GetView(); err == nil {
		v.Controller()
	}
	return 0
}

func (c *Init) Help() string {
	return ""
}

func (c *Init) Synopsis() string {
	return ""
}
