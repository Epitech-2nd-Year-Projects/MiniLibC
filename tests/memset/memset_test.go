package memset

import "testing"

func TestMemset(t *testing.T) {
	tests := []struct {
		name  string
		size  int
		value byte
	}{
		{name: "size 0", size: 0, value: 0},
		{name: "fill with 0", size: 10, value: 0},
		{name: "fill with 255", size: 10, value: 255},
		{name: "remplissage avec 42", size: 15, value: 42},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := make([]byte, tc.size)
			for i := range buf {
				buf[i] = 1 
			}
			Memset(buf, tc.value)
			for i, b := range buf {
				if b != tc.value {
					t.Errorf("index %d: expected %d, result %d", i, tc.value, b)
				}
			}
		})
	}
}

