package cases

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	lowered := strings.ToLower(s)
	runes := []rune(lowered)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
