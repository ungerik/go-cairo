// package cairo contains constants for use in cairo
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// LineCap cairo_line_cap_t
type LineCap int

// LineCap constants
const (
	LineCapButt LineCap = iota
	LineCapRound
	LineCapSquare
)

// LineJoin cairo_line_cap_join_t
type LineJoin int

// LineJoin constants
const (
	LineJoinMiter LineJoin = iota
	LineJoinRound
	LineJoinBevel
)
