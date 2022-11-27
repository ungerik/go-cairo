// Package cairo wraps the c cairographics library.
package cairo

import "C"

// FillRule cairo_fill_rule_t
type FillRule int

// FillRule constants
const (
	FillRuleWinding FillRule = iota
	FillRuleEvenOdd
)
