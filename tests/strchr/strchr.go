package strchr

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// char *strchr(const char *s, int c);
import "C"
import "unsafe"

func strchr(s string, c int) string {
	sc := C.CString(s)
	defer C.free(unsafe.Pointer(sc))
	result := C.strchr(sc, C.int(c))
	var ret string
	if result != nil {
		ret = C.GoString(result)
	}
	return ret
}
