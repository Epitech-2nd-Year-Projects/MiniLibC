package strlen

// #cgo LDFLAGS: -L../.. -lasm
//
// size_t strlen(const char *s);
import "C"

func strlen(str string) uint {
    strc := C.CString(str)
    return uint(C.strlen(strc))
}
