// Package cairo wraps the c cairographics library.
package cairo

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
