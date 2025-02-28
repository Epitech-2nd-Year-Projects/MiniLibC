package strrchr

import (
	"testing"
)

func TestStrrchr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		char     int
		expected string
	}{
		{name: "char at the end of the string", input: "hello", char: 'o', expected: "o"},
		{name: "char at the beginning", input: "hello", char: 'h', expected: "hello"},
		{name: "multiple char", input: "hello world", char: 'l', expected: "ld"},
		{name: "non present char", input: "hello", char: 'z', expected: ""},
		{name: "empty string", input: "", char: 'a', expected: ""},
		{name: "special char", input: "abc123#@", char: '#', expected: "#@"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := strrchr(tc.input, tc.char)
			if result != tc.expected {
				t.Errorf("strrchr(%s, %c) = %s, expected %s", tc.input, tc.char, result, tc.expected)
			}
		})
	}
}

