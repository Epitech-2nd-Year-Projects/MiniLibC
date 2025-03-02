package index

// #cgo LDFLAGS: -L../.. -lasm
// #include <strings.h>
// #include <stdlib.h>
//
// char *index(const char *s, int c);
import "C"
import "unsafe"

func index(s string, c byte) string {
	sc := C.CString(s)
	defer C.free(unsafe.Pointer(sc))
	res := C.index(sc, C.int(c))
	if res == nil {
		return ""
	}
	return C.GoString(res)
}

