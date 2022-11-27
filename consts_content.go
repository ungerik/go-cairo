// package cairo contains constants for use in cairo
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// Content cairo_content_t
type Content int

// Content constants
const (
	ContentColor      Content = 0x1000
	ContentAlpha      Content = 0x2000
	ContentColorAlpha Content = 0x3000
)
