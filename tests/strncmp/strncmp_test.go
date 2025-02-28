package strncmp

import "testing"

func TestStrncmp(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		n        int
		expected int 
	}{
		{"both empty", "", "", 5, 0},
		{"same strings", "hello", "hello", 5, 0},
		{"n < length | equal", "hello", "helium", 3, 0},
		{"n < length | difference", "hello", "helium", 4, int('l' - 'i')},
		{"s1 < s2", "apple", "apricot", 3, int('p' - 'r')},
		{"s1 > s2", "banana", "ban", 6, int('a' - 0)},
		{"n = 0", "any", "thing", 0, 0},
		{"n > length", "test", "test", 10, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Strncmp(tc.s1, tc.s2, tc.n)
			if result != tc.expected {
				t.Errorf("strncmp(%s, %s, %d) = %d, expected %d", tc.s1, tc.s2, tc.n, result, tc.expected)
			}
		})
	}
}

