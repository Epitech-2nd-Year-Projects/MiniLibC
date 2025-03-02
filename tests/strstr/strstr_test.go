package strstr

import "testing"

func TestStrstr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		expected string
	}{
		{name: "needle at the start", haystack: "hello world", needle: "hello", expected: "hello world"},
		{name: "needle in the middle", haystack: "hello world", needle: "lo wo", expected: "lo world"},
		{name: "needle at the end", haystack: "hello world", needle: "world", expected: "world"},
		{name: "needle not found", haystack: "hello world", needle: "test", expected: ""},
		{name: "empty needle", haystack: "hello world", needle: "", expected: "hello world"}, 
		{name: "empty haystack", haystack: "", needle: "hello", expected: ""},
		{name: "both empty", haystack: "", needle: "", expected: ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := strstr(tc.haystack, tc.needle)
			if result != tc.expected {
				t.Errorf("strstr(%s, %s) = %s, expected %s", tc.haystack, tc.needle, result, tc.expected)
			}
		})
	}
}

