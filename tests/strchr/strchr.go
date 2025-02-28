package strchr

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// char *strchr(const char *s, int c);
import "C"
import "unsafe"

func strchr(s string, c int) string {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result := C.strchr(cs, C.int(c))
	var ret string
	if result != nil {
		ret = C.GoString(result)
	}
	return ret
}
