package numbers

import "testing"

func TestHexToDec(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1E", "30"},
		{"42", "66"},
		{"0", "0"},
		{"a", "10"},
	}

	for _, tc := range tests {
		result, err := HexToDec(tc.input)

		if err != nil {
			t.Errorf("HexToDec(%q) вернул непредвиденную ошибку: %v", tc.input, err)
		}

		if result != tc.expected {
			t.Errorf("HexToDec(%q) = %q; ожидалось %q", tc.input, result, tc.expected)
		}
	}
}

func TestBinToDec(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"10", "2"},
		{"1111", "15"},
		{"0", "0"},
	}

	for _, tc := range tests {
		result, err := BinToDec(tc.input)

		if err != nil {
			t.Errorf("BinToDec(%q) вернул непредвиденную ошибку: %v", tc.input, err)
		}

		if result != tc.expected {
			t.Errorf("BinToDec(%q) = %q; ожидалось %q", tc.input, result, tc.expected)
		}
	}
}
