package strcmp

import (
	"testing"
)

func TestStrcmp(t *testing.T) {
    tests := []struct {
		name string
        s1 string
        s2 string
        expected int
    }{
		{name: "basic string", s1: "coucou", s2: "coucou", expected: 0},
		{name: "basic string with maj", s1: "abcd", s2: "abCd", expected: int('c' - 'C')},
		{name: "empty string", s1: "", s2: "", expected: 0},
		{name: "negative result", s1: "abCd", s2: "abcd", expected: int('C' - 'c')},
    }
    for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
        	result := strcmp(tc.s1, tc.s2)
        	if result != tc.expected {
            	t.Errorf("strcmp(%s, %s) = %d, expected %d", tc.s1, tc.s2, result, tc.expected)
        	}
		})
	}
}
