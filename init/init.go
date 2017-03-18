package init

import "github.com/ieee0824/thor/view"

func init() {
	view.View["JP"] = view.New()
	generator := SelectGeneratorType{}
	generator.Init()
	view.View["JP"].Add(generator, "selecGeneratorTypeView")

	askELB := AskUseELB{}
	askELB.Init()
	view.View["JP"].Add(askELB, "askUseELB")

	setELBName := ELBNameSetting{}
	setELBName.Init()
	view.View["JP"].Add(setELBName, "elbNameSetting")
}
