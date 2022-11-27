// Package cairo wraps the c cairographics library.
package cairo

// PathDataType cairo_path_data_type_t
type PathDataType int

// PathDataType constants
const (
	PathMoveTo PathDataType = iota
	PathLineTo
	PathCurveTo
	PathClosePath
)
