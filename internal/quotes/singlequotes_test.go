package quotes

import (
	"testing"
)

func TestSingleQuotes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Single word in quotes",
			input:    "I am exactly how they describe me: ' awesome '",
			expected: "I am exactly how they describe me: 'awesome'",
		},
		{
			name:     "Sentence in quotes",
			input:    "As Elton John said: ' I am the most well-known homosexual in the world '",
			expected: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			name:     "No quotes in text",
			input:    "Hello world without quotes",
			expected: "Hello world without quotes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SingleQuotes(tt.input)
			if got != tt.expected {
				t.Errorf("SingleQuotes(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}
