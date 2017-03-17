package init

import "github.com/ieee0824/thor/view"

func init() {
	view.View["JP"] = view.New()
	generator := SelectGeneratorType{}
	generator.Init()
	view.View["JP"].Add(generator, "selecGeneratorTypeView")

	creater := AskCreateELB{}
	creater.Init()
	view.View["JP"].Add(creater, "askCreateELB")
}
