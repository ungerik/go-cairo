// package cairo contains constants for use in cairo
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// FillRule cairo_fill_rule_t
type FillRule int

// FillRule constants
const (
	FillRuleWinding FillRule = iota
	FillRuleEvenOdd
)
