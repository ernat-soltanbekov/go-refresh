package grammar

import (
	"reflect"
	"testing"
)

func TestFixArticles(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"a", "apple"}, []string{"an", "apple"}},
		{[]string{"A", "orange"}, []string{"An", "orange"}},
		{[]string{"a", "rock"}, []string{"a", "rock"}},
		{[]string{"I", "found", "a", "amazing", "rock"}, []string{"I", "found", "an", "amazing", "rock"}},
	}
	for _, tc := range tests {
		inputCopy := append([]string(nil), tc.input...)
		result := FixArticles(inputCopy)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("FixArticles(%v) = %v; ожидалось %v", tc.input, result, tc.expected)
		}
	}
}
