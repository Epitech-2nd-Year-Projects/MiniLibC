package strrchr

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// char *strrchr(const char *s, int c);
import "C"
import "unsafe"

func strrchr(s string, c int) string {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result := C.strrchr(cs, C.int(c))
	var ret string
	if result != nil {
		ret = C.GoString(result)
	}
	return ret
}
