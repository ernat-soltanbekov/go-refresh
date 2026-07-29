package punct

import (
	"regexp"
	"strings"
)

func FixSpacing(words []string) []string {
	if len(words) == 0 {
		return nil
	}
	line := strings.Join(words, " ")
	reBefore := regexp.MustCompile(`\s+([.,!?:;]+)`)
	line = reBefore.ReplaceAllString(line, "$1")
	line = strings.ReplaceAll(line, "!?...", "!? ...")
	line = strings.ReplaceAll(line, "!...", "! ...")
	line = strings.ReplaceAll(line, "?...", "? ...")
	line = strings.ReplaceAll(line, ":...", ": ...")
	reAfter := regexp.MustCompile(`([.,!?:;]+)([\pL0-9])`)
	line = reAfter.ReplaceAllString(line, "$1$ 2")
	return strings.Fields(line)
}
