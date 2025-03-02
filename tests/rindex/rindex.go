package rindex

// #cgo LDFLAGS: -L../.. -lasm
// #include <strings.h>
// #include <stdlib.h>
//
// char *rindex(const char *s, int c);
import "C"
import "unsafe"

func rindex(s string, c byte) string {
	sc := C.CString(s)
	defer C.free(unsafe.Pointer(sc))
	res := C.rindex(sc, C.int(c))
	if res == nil {
		return ""
	}
	return C.GoString(res)
}

