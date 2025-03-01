package strpbrk

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// char *strpbrk(const char *s, const char *accept);
import "C"
import "unsafe"

func strpbrk(s, accept string) string {
	cs := C.CString(s)
	caccept := C.CString(accept)
	defer C.free(unsafe.Pointer(cs))
	defer C.free(unsafe.Pointer(caccept))
	
	result := C.strpbrk(cs, caccept)
	if result == nil {
		return ""
	}
	return C.GoString(result)
}

