package strncmp

// #cgo LDFLAGS: -L../.. -lasm
// #include <stdlib.h>
//
// int strncmp(const char *s1, const char *s2, size_t n);
import "C"
import "unsafe"

func Strncmp(s1, s2 string, n int) int {
    sc1 := C.CString(s1)
    sc2 := C.CString(s2)
    defer C.free(unsafe.Pointer(sc1))  // Libère la mémoire allouée
    defer C.free(unsafe.Pointer(sc2))  // Après l'appel à C.strncmp
    return int(C.strncmp(sc1, sc2, C.size_t(n)))
}
