package punct

import (
	"reflect"
	"testing"
)

func TestFixSpacing(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Базовая пунктуация",
			input:    []string{"Hello", ",", "world", "."},
			expected: []string{"Hello,", "world."},
		},
		{
			name:     "Все виды знаков препинания.",
			input:    []string{"a", ",", "b", ".", "c", "!", "d", "?", "e", ":", "f", ";"},
			expected: []string{"a,", "b.", "c!", "d?", "e:", "f;"},
		},
		{
			name:     "Несколько знаков препинания подряд",
			input:    []string{"wait", "!", "?", "..."},
			expected: []string{"wait!?", "..."},
		},
		{
			name:     "Обычные слова без знаков препинания",
			input:    []string{"just", "some", "words"},
			expected: []string{"just", "some", "words"},
		},
		{
			name:     "Пустой массив на входе",
			input:    []string{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FixSpacing(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FixSpacing(%v) = %v; ожидалось %v", tt.input, result, tt.expected)
			}
		})
	}
}
