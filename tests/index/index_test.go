package index

import "testing"

func TestIndex(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		c        byte
		expected string
	}{
		{name: "caractère au début", s: "hello", c: 'h', expected: "hello"},
		{name: "caractère au milieu", s: "hello", c: 'l', expected: "llo"},
		{name: "caractère à la fin", s: "hello", c: 'o', expected: "o"},
		{name: "caractère non trouvé", s: "hello", c: 'z', expected: ""},
		{name: "chaîne vide", s: "", c: 'a', expected: ""},
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

