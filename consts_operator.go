// package cairo contains constants for use in cairo
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// Operator cairo_operator_t
type Operator int

// Operator constants
const (
	OperatorClear = iota

	OperatorSource
	OperatorOver
	OperatorIn
	OperatorOut
	OperatorAtop

	OperatorDest
	OperatorDestOver
	OperatorDestIn
	OperatorDestOut
	OperatorDestAtop

	OperatorXor
	OperatorAdd
	OperatorSaturate

	OperatorMultiply
	OperatorScreen
	OperatorOverlay
	OperatorDarken
	OperatorLighten
	OperatorColorDodge
	OperatorColorBurn
	OperatorHardLight
	OperatorSoftLight
	OperatorDifference
	OperatorExclusion
	OperatorHslHue
	OperatorHslSaturation
	OperatorHslColor
	OperatorHslLuminosity
)
