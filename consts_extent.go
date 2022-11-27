// Package cairo wraps the c cairographics library.
package cairo

// Extent cairo_extent_t
type Extentx int

// Extent constants
const (
	ExtendNone Extentx = iota
	ExtendRepeat
	ExtendReflect
	ExtendPad
)
