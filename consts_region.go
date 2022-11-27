// Package cairo wraps the c cairographics library.
package cairo

// RegionOverlap Used as the return value for cairo_region_contains_rectangle().
type RegionOverlap int

// RegionOverlap constants
const (
	RegionOverlapIn RegionOverlap = iota
	RegionOverlapOut
	RegionOverlapPart
)
