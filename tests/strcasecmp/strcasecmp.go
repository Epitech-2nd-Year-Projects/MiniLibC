package strcasecmp

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// int strcasecmp(const char *s1, const char *s2);
import "C"
import "unsafe"

func strcasecmp(s1, s2 string) int {
	sc1 := C.CString(s1)
	sc2 := C.CString(s2)
	defer C.free(unsafe.Pointer(sc1))
	defer C.free(unsafe.Pointer(sc2))
	return int(C.strcasecmp(sc1, sc2))
}

