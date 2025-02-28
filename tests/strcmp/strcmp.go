package strcmp

// #cgo LDFLAGS: -L../.. -lasm
//
// int strcmp(const char *str1, const char *str2);
import "C"

func strcmp(str1 string, str2 string) C.int {
    strc1 := C.CString(str1)
    strc2 := C.CString(str2)
    result := C.strcmp(strc1, strc2)
    return result
}

