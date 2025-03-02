package ffs

// #cgo LDFLAGS: -L../.. -lasm
// #include <strings.h>
import "C"

func ffs(i int) int {
	return int(C.ffs(C.int(i)))
}

