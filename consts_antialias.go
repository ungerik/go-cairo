// package cairo contains constants for use in cairo
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// Antialias cairo_antialias_t
type Antialias int

// Antialias constants
const (
	AntialiasDefault Antialias = iota
	AntialiasNone
	AntialiasGray
	AntialiasSubpixel
)
