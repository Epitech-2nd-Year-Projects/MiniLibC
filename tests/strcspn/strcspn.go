package strcspn

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// size_t strcspn(const char *s, const char *reject);
import "C"
import "unsafe"

func strcspn(s, reject string) int {
	sc := C.CString(s)
	creject := C.CString(reject)
	defer C.free(unsafe.Pointer(sc))
	defer C.free(unsafe.Pointer(creject))
	return int(C.strcspn(sc, creject))
}

