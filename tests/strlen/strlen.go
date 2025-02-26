package strlen

// #cgo LDFLAGS: -L../.. -lasm
//
// size_t strlen(const char *s);
import "C"

func strlen(str string) uint64 {
    strc := C.CString(str)
    result := C.strlen(strc)
    return uint64(result)
}
