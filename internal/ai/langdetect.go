package ai

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func LangDetect(text string) string {
	lowerText := strings.ToLower(text)
	engWords := []string{"the", "and", "of", "to", "is", "in", "you", "it", "that"}
	frWords := []string{"le", "la", "les", "de", "et", "un", "une", "est", "dans", "vous"}
	accentedChars := []string{"é", "è", "à", "ç", "ô", "î", "û"}
	engCount := 0
	frCount := 0
	reWord := regexp.MustCompile(`[a-zA-ZàâäéèêëîïôöùûüçÀÂÄÉÈÊËÎÏÔÖÙÛÜÇ]+`)
	words := reWord.FindAllString(lowerText, -1)
	engMap := make(map[string]bool)
	for _, w := range engWords {
		engMap[w] = true
	}
	frMap := make(map[string]bool)
	for _, w := range frWords {
		frMap[w] = true
	}
	for _, w := range words {
		if engMap[w] {
			engCount++
		}
		if frMap[w] {
			frCount++
		}
	}
	for _, char := range accentedChars {
		frCount += strings.Count(lowerText, char)
	}
	total := engCount + frCount
	if total == 0 {
		return "Language: Unknown"
	}
	if engCount >= frCount {
		pct := int(math.Round(float64(engCount) * 100.0 / float64(total)))
		return fmt.Sprintf("Language: English (%d%%)", pct)
	} else {
		pct := int(math.Round(float64(frCount) * 100.0 / float64(total)))
		return fmt.Sprintf("Language: French (%d%%)", pct)
	}
}
