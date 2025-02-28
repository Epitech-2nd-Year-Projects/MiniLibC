package strcasecmp

import (
	"testing"
)

func TestStrcasecmp(t *testing.T) {
	tests := []struct {
		name      string
		s1        string
		s2        string
		expected int 
	}{
		{name: "same string", s1: "hello", s2: "hello", expected: 0},
		{name: "same string with different case", s1: "hello", s2: "HELLO", expected: 0},
		{name: "s1 < s2", s1: "abc", s2: "abd", expected: -1},
		{name: "s1 > s2", s1: "ABD", s2: "ABC", expected: 1},
		{name: "fully different", s1: "foo", s2: "bar", expected: 4},
		{name: "empty string", s1: "", s2: "", expected: 0},
		{name: "empty s1", s1: "", s2: "aaa", expected: -97},
		{name: "empty s2", s1: "aaa", s2: "", expected: 97},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := strcasecmp(tc.s1, tc.s2)
			if result != tc.expected {
				t.Errorf("strcasecmp(%s, %s) = %d, expected %d", tc.s1, tc.s2, result, tc.expected)
			}
		})
	}
}

