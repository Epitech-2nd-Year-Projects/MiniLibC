package strcmp

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// int strcmp(const char *str1, const char *str2);
import "C"
import "unsafe"

func strcmp(s1, s2 string) int {
    sc1 := C.CString(s1)
    sc2 := C.CString(s2)
	defer C.free(unsafe.Pointer(sc1))
	defer C.free(unsafe.Pointer(sc2))
    result := C.strcmp(sc1, sc2)
    return int(result)
}

