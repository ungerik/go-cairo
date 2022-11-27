// Package cairo wraps the c cairographics library.
package cairo

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
