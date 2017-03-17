package util

import (
	"errors"
	"strings"
)

type LANG string

func IsMultiByteChar(r rune) bool {
	c := string([]rune{r})
	return 1 != len(c)
}

func ConvertRune(s string) (rune, error) {
	if len([]rune(s)) != 1 {
		var e rune
		return e, errors.New("not char")
	}
	return []rune(s)[0], nil
}

func MultiString(s string) string {
	array := strings.Split(s, "")
	var ret string

	for _, s := range array {
		r, _ := ConvertRune(s)
		if IsMultiByteChar(r) {
			ret = ret + s + " "
		} else {
			ret += s
		}
	}
	return ret
}
