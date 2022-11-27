// Package cairo wraps the c cairographics library.
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// Version cairo_version
func Version() int {
	return int(C.cairo_version())
}

// VersionString cairo_version_string
func VersionString() string {
	return C.GoString(C.cairo_version_string())
}
