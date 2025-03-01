package strpbrk

import "testing"

func TestStrpbrk(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		accept   string
		expected string
	}{
		{"first char matching", "hello", "h", "hello"},
		{"middle matching", "hello", "e", "ello"},
		{"matching with multiple accept", "hello", "ol", "llo"},
		{"no match", "hello", "z", ""},
		{"empty string", "", "abc", ""},
		{"empty accept", "hello", "", ""},
		{"both empty", "", "", ""},
		{"complex matching", "abcdef", "xdy", "def"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := strpbrk(tc.s, tc.accept)
			if result != tc.expected {
				t.Errorf("strpbrk(%s, %s) = %s, expected %s", tc.s, tc.accept, result, tc.expected)
			}
		})
	}
}

