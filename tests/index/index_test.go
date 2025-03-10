package index

import "testing"

func TestIndex(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		c        byte
		expected string
	}{
		{name: "char at the beginning", s: "hello", c: 'h', expected: "hello"},
		{name: "char in the middle", s: "hello", c: 'l', expected: "llo"},
		{name: "char at the end", s: "hello", c: 'o', expected: "o"},
		{name: "char not found", s: "hello", c: 'z', expected: ""},
		{name: "empty string", s: "", c: 'a', expected: ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := index(tc.s, tc.c)
			if result != tc.expected {
				t.Errorf("index(%s, %s) = %s, expected %s", tc.s, string(tc.c), result, tc.expected)
			}
		})
	}
}

