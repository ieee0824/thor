package init

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/elbv2"
	c "github.com/ieee0824/thor/controller"
	"github.com/ieee0824/thor/elb"
	"github.com/ieee0824/thor/util"
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

func getLoadBalancers() ([]*elbv2.LoadBalancer, error) {
	input := &elbv2.DescribeLoadBalancersInput{}
	out, err := elb.DescribeLoadBalancers(util.AwsConfig, input)
	if err != nil {
		return nil, err
	}
	return out.LoadBalancers, nil
}

func isExistELB(name string, lbs []*elbv2.LoadBalancer) bool {
	for _, lb := range lbs {
		if *lb.LoadBalancerName == name {
			return true
		}
	}
	return false
}

func containELB(name string, lbs []*elbv2.LoadBalancer) []string {
	var ret []string
	for _, lb := range lbs {
		if strings.Contains(*lb.LoadBalancerName, name) {
			ret = append(ret, *lb.LoadBalancerName)
		}
	}
	return ret
}

func setELBNameController() {
	lbs, err := getLoadBalancers()
	if err != nil {
		panic(err)
	}
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
						if isExistELB(box.Answer(), lbs) {
							view.View["JP"].Fin()
							return
						}
					}
				}
			case termbox.KeyBackspace:
				if box, err := view.View["JP"].GetView(); err == nil {
					if b, ok := box.(ELBNameSetting); ok {
						b.BS()
						lines := containELB(box.Answer(), lbs)
						b.SetMeta(strings.Join(lines, "\n"))
					}
					c.Draw(box)
				}
			case termbox.KeyBackspace2:
				if box, err := view.View["JP"].GetView(); err == nil {
					if b, ok := box.(ELBNameSetting); ok {
						b.BS()
						lines := containELB(box.Answer(), lbs)
						b.SetMeta(strings.Join(lines, "\n"))
					}
					c.Draw(box)
				}
			default:
				if box, err := view.View["JP"].GetView(); err == nil {
					if b, ok := box.(ELBNameSetting); ok {
						b.Add(ev.Ch)
						lines := containELB(box.Answer(), lbs)
						b.SetMeta(strings.Join(lines, "\n"))
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
