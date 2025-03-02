package memcpy

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// void *memcpy(void *dest, const void *src, size_t n);
import "C"
import "unsafe"

func memcpy(dest, src []byte, n int) {
	if n > len(src) || n > len(dest) {
		panic("buffer too small to cpy n byte")
	}
	C.memcpy(unsafe.Pointer(&dest[0]), unsafe.Pointer(&src[0]), C.size_t(n))
}

