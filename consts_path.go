// package cairo contains constants for use in cairo
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// PathDataType cairo_path_data_type_t
type PathDataType int

// PathDataType constants
const (
	PathMoveTo PathDataType = iota
	PathLineTo
	PathCurveTo
	PathClosePath
)
