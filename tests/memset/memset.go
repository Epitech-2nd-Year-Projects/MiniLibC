package memset

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// void *memset(void *s, int c, size_t n);
import "C"
import "unsafe"

func Memset(s []byte, c byte) {
	if len(s) == 0 {
		return
	}
	C.memset(unsafe.Pointer(&s[0]), C.int(c), C.size_t(len(s)))
}

