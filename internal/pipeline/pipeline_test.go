package pipeline

import (
	"reflect"
	"testing"
)

func TestProcessTags(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"hello", "(up)", "(cap)"}, []string{"Hello"}},
		{[]string{"hello", "(up,"}, []string{"hello", "(up,"}},
		{[]string{"hello", "(up)"}, []string{"HELLO"}},
		{[]string{"ready", "set", "go", "(up)"}, []string{"ready", "set", "GO"}},
		{[]string{"1E", "(hex)", "files"}, []string{"30", "files"}},
		{[]string{"10", "(bin)", "years"}, []string{"2", "years"}},
		{[]string{"fOOLISH", "(cap)"}, []string{"Foolish"}},
		{[]string{"обычное", "предложение"}, []string{"обычное", "предложение"}},
	}
	for _, tc := range tests {
		inputCopy := append([]string(nil), tc.input...)
		result := ProcessTags(inputCopy)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("ProcessTags(%v) = %v; ожидалось %v", tc.input, result, tc.expected)
		}
	}
}
