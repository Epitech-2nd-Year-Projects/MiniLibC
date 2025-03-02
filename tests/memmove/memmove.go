package memmove

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// void *memmove(void *dest, const void *src, size_t n);
import "C"
import "unsafe"

func memmove(dest, src []byte, n int) {
	if n > len(src) || n > len(dest) {
		panic("buffer too small to cpy n byte")
	}
	C.memmove(unsafe.Pointer(&dest[0]), unsafe.Pointer(&src[0]), C.size_t(n))
}

