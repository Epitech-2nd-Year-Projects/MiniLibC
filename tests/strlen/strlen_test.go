package strlen

import (
    "testing"
)

func TestStrlen(t *testing.T) {
    tests := []struct {
		name string
        str string
        expected uint
    }{
		{name: "simple string", str: "Hello", expected: 5},
		{name: "empty string", str: "", expected: 0},
		{name: "string with spaces", str: " Hello ", expected: 7},
		{name: "special characters", str: "abc123!@#", expected: 9},
    }
    for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := strlen(tc.str)
			if result != tc.expected {
				t.Errorf("strlen(%s) = %d, expected %d", tc.str, result, tc.expected)
			}
		})
    }
}
