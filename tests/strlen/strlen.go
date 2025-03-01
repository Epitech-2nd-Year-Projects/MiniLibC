package strlen

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// size_t strlen(const char *s);
import "C"
import "unsafe"

func strlen(str string) uint {
    sc := C.CString(str)
	defer C.free(unsafe.Pointer(sc))
    return uint(C.strlen(sc))
}
