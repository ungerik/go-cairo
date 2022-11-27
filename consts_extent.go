// package cairo contains constants for use in cairo
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// Extent cairo_extent_t
type Extent int

// Extent constants
const (
	ExtendNone Extent = iota
	ExtendRepeat
	ExtendReflect
	ExtendPad
)
