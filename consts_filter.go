// package cairo contains constants for use in cairo
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// Filter cairo_filter_t
type Filter int

// Filter constants
const (
	CairoFilterFast Filter = iota
	CairoFilterGood
	CairoFilterBest
	CairoFilterNearest
	CairoFilterBilinear
	CairoFilterGaussian
)
