package cases

import "testing"

func TestTransform(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "Hello"},
		{"WORLD", "World"},
		{"gO", "Go"},
		{"", ""},
		{"привет", "Привет"},
	}

	for _, tc := range tests {
		result := Capitalize(tc.input)
		if result != tc.expected {
			t.Errorf("Capitalize(%q) = %q ожидалось %q", tc.input, result, tc.expected)
		}
	}
}
