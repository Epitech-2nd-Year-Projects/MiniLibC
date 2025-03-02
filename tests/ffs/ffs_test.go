package ffs

import "testing"

func TestFfs(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{name: "0", input: 0, expected: 0},
		{name: "byte 0 to 1", input: 1, expected: 1},
		{name: "byte 1 to 1", input: 2, expected: 2},
		{name: "byte 0 and 1 to 1", input: 3, expected: 1},
		{name: "byte 4 to 1", input: 16, expected: 5},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ffs(tc.input)
			if result != tc.expected {
				t.Errorf("ffs(%d) = %d, expected %d", tc.input, result, tc.expected)
			}
		})
	}
}

