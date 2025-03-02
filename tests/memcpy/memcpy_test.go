package memcpy

import (
	"bytes"
	"testing"
)

func TestMemcpy(t *testing.T) {
	tests := []struct {
		name     string
		src      []byte
		dest     []byte
		n        int
		expected []byte
	}{
		{name: "full copy", src: []byte{1, 2, 3, 4}, dest: []byte{0, 0, 0, 0}, n: 4, expected: []byte{1, 2, 3, 4}},
		{name: "partial copy", src: []byte{10, 20, 30, 40}, dest: []byte{5, 5, 5, 5}, n: 2, expected: []byte{10, 20, 5, 5}},
		{name: "0 byte copy", src: []byte{1, 2, 3, 4}, dest: []byte{9, 9, 9, 9}, n: 0, expected: []byte{9, 9, 9, 9}},
		{name: "copy on a larger array", src: []byte{7, 8, 9}, dest: []byte{0, 0, 0, 0, 0}, n: 3, expected: []byte{7, 8, 9, 0, 0}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dst := make([]byte, len(tc.dest))
			copy(dst, tc.dest)
			memcpy(dst, tc.src, tc.n)
			if !bytes.Equal(dst, tc.expected) {
				t.Errorf("memcpy() = %v, expected %v", dst, tc.expected)
			}
		})
	}
}

