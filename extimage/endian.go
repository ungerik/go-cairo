package extimage

import (
	"unsafe"
)

var (
	tst          uint32 = 1
	littleEndian bool   = (*[4]byte)(unsafe.Pointer(&tst))[0] == 1
)
