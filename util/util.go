package util

import "strings"

type LANG string

func MultiString(s string) string {
	array := strings.Split(s, "")
	return strings.Join(array, " ")
}
