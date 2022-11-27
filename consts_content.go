// Package cairo wraps the c cairographics library.
package cairo

// Content cairo_content_t
type Content int

// Content constants
const (
	ContentColor      Content = 0x1000
	ContentAlpha      Content = 0x2000
	ContentColorAlpha Content = 0x3000
)
