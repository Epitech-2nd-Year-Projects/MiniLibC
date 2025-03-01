package strcspn

import "testing"

func TestStrcspn(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		reject   string
		expected int
	}{
		{name: "no reject", s: "hello", reject: "z", expected: 5},
		{name: "reject at the beginning", s: "hello", reject: "h", expected: 0},
		{name: "reject in the middle", s: "hello", reject: "l", expected: 2},
		{name: "multiple reject", s: "abcdef", reject: "dx", expected: 3},
		{name: "empty string", s: "", reject: "abc", expected: 0},
		{name: "empty reject", s: "hello", reject: "", expected: 5},
		{name: "both empty", s: "", reject: "", expected: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := strcspn(tc.s, tc.reject)
			if result != tc.expected {
				t.Errorf("strcspn(%s, %s) = %d, attendu %d", tc.s, tc.reject, result, tc.expected)
			}
		})
	}
}

