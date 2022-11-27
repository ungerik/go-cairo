// Package cairo wraps the c cairographics library.
package cairo

// Antialias cairo_antialias_t
type Antialias int

// Antialias constants
const (
	AntialiasDefault Antialias = iota
	AntialiasNone
	AntialiasGray
	AntialiasSubpixel
)
