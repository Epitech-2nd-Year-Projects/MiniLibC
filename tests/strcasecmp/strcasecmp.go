package strcasecmp

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// int strcasecmp(const char *s1, const char *s2);
import "C"
import "unsafe"

func strcasecmp(s1, s2 string) int {
	cs1 := C.CString(s1)
	cs2 := C.CString(s2)
	defer C.free(unsafe.Pointer(cs1))
	defer C.free(unsafe.Pointer(cs2))
	return int(C.strcasecmp(cs1, cs2))
}

