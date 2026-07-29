package pipeline

import (
	"go-refresh/internal/cases"
	"go-refresh/internal/grammar"
	"go-refresh/internal/numbers"
	"go-refresh/internal/punct"
	"go-refresh/internal/quotes"
	"regexp"
	"strconv"
	"strings"
)

func Process(text string) string {
	lines := strings.Split(text, "\n")
	var processedLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			processedLines = append(processedLines, "")
			continue
		}
		words := strings.Fields(line)
		words = ProcessTags(words)
		lineText := strings.Join(words, " ")
		lineText = quotes.SingleQuotes(lineText)
		words = strings.Fields(lineText)
		words = punct.FixSpacing(words)
		words = grammar.FixArticles(words)
		processedLines = append(processedLines, strings.Join(words, " "))
	}
	return strings.Join(processedLines, "\n")
}

func isValidWord(token string) bool {
	matched, _ := regexp.MatchString(`[\pL0-9]`, token)
	return matched
}

func ProcessTags(words []string) []string {
	var result []string
	for i := 0; i < len(words); i++ {
		word := words[i]
		switch word {
		case "(up)":
			applyTransform(&result, 1, cases.ToUpper)
		case "(low)":
			applyTransform(&result, 1, cases.ToLower)
		case "(cap)":
			applyTransform(&result, 1, cases.Capitalize)
		case "(hex)":
			if len(result) > 0 {
				idx := getLastWordIndex(result)
				if idx >= 0 {
					dec, err := numbers.HexToDec(result[idx])
					if err == nil {
						result[idx] = dec
					}
				}
			}
		case "(bin)":
			if len(result) > 0 {
				idx := getLastWordIndex(result)
				if idx >= 0 {
					dec, err := numbers.BinToDec(result[idx])
					if err == nil {
						result[idx] = dec
					}
				}
			}
		case "(up,", "(low,", "(cap,":
			if i+1 < len(words) {
				numStr := strings.TrimSuffix(words[i+1], ")")
				n, err := strconv.Atoi(numStr)
				if err == nil {
					switch word {
					case "(up,":
						applyTransform(&result, n, cases.ToUpper)
					case "(low,":
						applyTransform(&result, n, cases.ToLower)
					case "(cap,":
						applyTransform(&result, n, cases.Capitalize)
					}
					i++
					continue
				}
			}
			result = append(result, word)
		default:
			if strings.HasPrefix(word, "(up,") || strings.HasPrefix(word, "(low,") || strings.HasPrefix(word, "(cap,") {
				cleaned := strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(strings.TrimPrefix(word, "(up,"), "(low,"), "(cap,"), ")")
				cleaned = strings.TrimSpace(cleaned)
				n, err := strconv.Atoi(cleaned)
				if err == nil {
					fn := cases.ToUpper
					if strings.HasPrefix(word, "(low,") {
						fn = cases.ToLower
					} else if strings.HasPrefix(word, "(cap,") {
						fn = cases.Capitalize
					}
					applyTransform(&result, n, fn)
					continue
				}
			}
			result = append(result, word)
		}
	}
	return result
}

func getLastWordIndex(res []string) int {
	for j := len(res) - 1; j >= 0; j-- {
		if isValidWord(res[j]) {
			return j
		}
	}
	return -1
}

func applyTransform(res *[]string, n int, transform func(string) string) {
	words := *res
	count := 0
	for j := len(words) - 1; j >= 0 && count < n; j-- {
		if isValidWord(words[j]) {
			words[j] = transform(words[j])
			count++
		}
	}
}
