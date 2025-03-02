package memmove

import (
	"bytes"
	"testing"
)

func TestMemmoveNonOverlap(t *testing.T) {
	tests := []struct {
		name     string
		src      []byte
		dest     []byte
		n        int
		expected []byte
	}{
		{name: "full cpy", src: []byte{1, 2, 3, 4}, dest: []byte{0, 0, 0, 0}, n: 4, expected: []byte{1, 2, 3, 4}},
		{name: "partial cpy", src: []byte{10, 20, 30, 40}, dest: []byte{5, 5, 5, 5}, n: 2, expected: []byte{10, 20, 5, 5}},
		{name: "cpy of 0 byte", src: []byte{1, 2, 3, 4}, dest: []byte{9, 9, 9, 9}, n: 0, expected: []byte{9, 9, 9, 9}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dst := make([]byte, len(tc.dest))
			copy(dst, tc.dest)
			memmove(dst, tc.src, tc.n)
			if !bytes.Equal(dst, tc.expected) {
				t.Errorf("memmove() = %v, expected %v", dst, tc.expected)
			}
		})
	}
}
