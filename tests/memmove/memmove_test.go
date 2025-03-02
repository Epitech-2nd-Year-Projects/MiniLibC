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
		{name: "copie complète sans chevauchement", src: []byte{1, 2, 3, 4}, dest: []byte{0, 0, 0, 0}, n: 4, expected: []byte{1, 2, 3, 4}},
		{name: "copie partielle sans chevauchement", src: []byte{10, 20, 30, 40}, dest: []byte{5, 5, 5, 5}, n: 2, expected: []byte{10, 20, 5, 5}},
		{name: "copie de zéro octet", src: []byte{1, 2, 3, 4}, dest: []byte{9, 9, 9, 9}, n: 0, expected: []byte{9, 9, 9, 9}},
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

func TestMemmoveOverlap(t *testing.T) {
	// Test avec zones se chevauchant dans le même tableau
	buf := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	// On souhaite copier les 4 premiers octets à partir de l'index 0 vers l'index 2.
	// Avant copie : [1,2,3,4,5,6,7,8]
	// Après copie attendue : [1,2,1,2,3,4,7,8]
	expected := []byte{1, 2, 1, 2, 3, 4, 7, 8}
	memmove(buf[2:], buf, 4)
	if !bytes.Equal(buf, expected) {
		t.Errorf("Memmove overlap = %v, attendu %v", buf, expected)
	}
}

func TestMemmoveIdentiqueSourceDestination(t *testing.T) {
	// Cas où src et dest sont la même slice
	buf := []byte{9, 9, 9, 9}
	expected := []byte{9, 9, 9, 9}
	memmove(buf, buf, 4)
	if !bytes.Equal(buf, expected) {
		t.Errorf("Memmove identique = %v, attendu %v", buf, expected)
	}
}

