package grammar

import (
	"regexp"
	"strings"
)

func FixArticles(words []string) []string {
	text := strings.Join(words, " ")
	reLower := regexp.MustCompile(`\ba\s+([aeiouAEIOUhH])`)
	text = reLower.ReplaceAllString(text, "an $1")
	reUpper := regexp.MustCompile(`\bA\s+([aeiouAEIOUhH])`)
	text = reUpper.ReplaceAllString(text, "An $1")
	return strings.Fields(text)
}
