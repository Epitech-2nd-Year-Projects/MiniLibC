package strlen

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// size_t strlen(const char *s);
import "C"
import "unsafe"

func strlen(str string) uint {
    strc := C.CString(str)
	defer C.free(unsafe.Pointer(strc))
    return uint(C.strlen(strc))
}
