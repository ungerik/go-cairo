// package cairo contains constants for use in cairo
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// RegionOverlap Used as the return value for cairo_region_contains_rectangle().
type RegionOverlap int

// RegionOverlap constants
const (
	RegionOverlapIn RegionOverlap = iota
	RegionOverlapOut
	RegionOverlapPart
)
