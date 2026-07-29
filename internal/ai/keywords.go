package ai

import (
	"regexp"
	"sort"
	"strings"
)

func ExtractKeywords(text string) string {
	text = strings.ToLower(text)
	re := regexp.MustCompile(`[a-z0-9']+`)
	words := re.FindAllString(text, -1)
	stopwordsList := []string{
		"the", "and", "of", "to", "is", "in", "it", "that", "a", "an", "for", "on",
		"with", "as", "by", "at", "from", "this", "these", "those", "de", "la",
		"le", "les", "et", "un", "une", "des",
	}
	stopwords := make(map[string]bool)
	for _, word := range stopwordsList {
		stopwords[word] = true
	}
	counts := make(map[string]int)
	for _, w := range words {
		if !stopwords[w] {
			counts[w]++
		}
	}
	if len(counts) == 0 {
		return "Keywords: (none)"
	}
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range counts {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Key < ss[j].Key
		}
		return ss[i].Value > ss[j].Value
	})
	limit := 5
	if len(ss) < limit {
		limit = len(ss)
	}
	var result []string
	for i := 0; i < limit; i++ {
		result = append(result, ss[i].Key)
	}
	return "Keywords: " + strings.Join(result, ", ")
}
