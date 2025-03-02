package strstr

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// char *strstr(const char *haystack, const char *needle);
import "C"
import "unsafe"

func strstr(haystack, needle string) string {
	sc := C.CString(haystack)
	cneedle := C.CString(needle)
	defer C.free(unsafe.Pointer(sc))
	defer C.free(unsafe.Pointer(cneedle))
	result := C.strstr(sc, cneedle)
	if result == nil {
		return ""
	}
	return C.GoString(result)
}

