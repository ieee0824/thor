package init

var JPMenu = map[string]Menu{
	"selecGeneratorType": selecGeneratorTypeJP(),
}

func selecGeneratorTypeJP() *SelectBox {
	return NewSelectBox(
		"設定ファイルの生成方法を選ぶ",
		[]string{
			"対話式に設定をする",
			"テンプレートを生成する",
		},
	)
}
