// Package cairo wraps the c cairographics library.
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// Format cairo_format_t
type Format int

// StrideForWidth provides a stride value that will respect all alignment requirements of the accelerated image-rendering code within cairo
func (f Format) StrideForWidth(width int) int {
	return int(C.cairo_format_stride_for_width(C.cairo_format_t(f), C.int(width)))
}

// Format constants
const (
	FormatInvalid  Format = -1
	FormatARGB32   Format = 0
	FormatRGB24    Format = 1
	FormatA8       Format = 2
	FormatA1       Format = 3
	FormatRGB16565 Format = 4
	FormatRGB30    Format = 5
)
