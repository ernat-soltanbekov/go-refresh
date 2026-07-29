package quotes

import (
	"regexp"
)

func SingleQuotes(text string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return re.ReplaceAllString(text, "'$1'")
}
