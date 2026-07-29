package ai

import "testing"

func TestLangDetect(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "100% English",
			input:    "the and of to",
			expected: "Language: English (100%)",
		},
		{
			name:     "100% French with accents",
			input:    "le la les déjà",
			expected: "Language: French (100%)",
		},
		{
			name:     "Tie - English wins",
			input:    "the le",
			expected: "Language: English (50%)",
		},
		{
			name:     "Unknown language",
			input:    "hello world",
			expected: "Language: Unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LangDetect(tt.input)
			if got != tt.expected {
				t.Errorf("LangDetect() = %v, want %v", got, tt.expected)
			}
		})
	}
}
