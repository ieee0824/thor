package init

type menuName string

var transitions = []menuName{
	"generateQuestinType",
}

type Menu interface {
	Answer() string
	Message() string
	ToggleCursor()
}
