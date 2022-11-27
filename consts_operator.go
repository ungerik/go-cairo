// Package cairo wraps the c cairographics library.
package cairo

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
