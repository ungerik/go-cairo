// +build !goci

package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// PatternType represents a cairo_pattern_type_t
type PatternType int

// Pattern types
const (
	PatternTypeSolid PatternType = iota
	PatternTypeSurface
	PatternTypeLinear
	PatternTypeRadial
)

// Pattern represents a cairo_pattern_t
type Pattern struct {
	pattern *C.cairo_pattern_t
}

////////////////////////////
// pattern creation

// CreateLinearGradient creates a pattern to be used as a linear gradient.
func CreateLinearGradient(x0, y0, x1, y1 float64) *Pattern {
	p := C.cairo_pattern_create_linear(C.double(x0), C.double(y0), C.double(x1), C.double(y1))
	return &Pattern{p}
}

// CreateRadialGradient creates a pattern to be used as a radial gradient.
func CreateRadialGradient(cx0, cy0, radius0, cx1, cy1, radius1 float64) *Pattern {
	p := C.cairo_pattern_create_radial(
		C.double(cx0), C.double(cy0), C.double(radius0),
		C.double(cx1), C.double(cy1), C.double(radius1),
	)
	return &Pattern{p}
}

// CreateRGBPattern creates a solid RGB pattern.
func CreateRGBPattern(red, green, blue float64) *Pattern {
	p := C.cairo_pattern_create_rgb(C.double(red), C.double(green), C.double(blue))
	return &Pattern{p}
}

// CreateRGBAPattern creates a solid RGBA pattern.
func CreateRGBAPattern(red, green, blue, alpha float64) *Pattern {
	p := C.cairo_pattern_create_rgba(C.double(red), C.double(green), C.double(blue), C.double(alpha))
	return &Pattern{p}
}

////////////////////////////
// gradient methods

// AddColorStopRGB adds an rgb color stop to this pattern.
func (p *Pattern) AddColorStopRGB(offset, red, green, blue float64) {
	C.cairo_pattern_add_color_stop_rgb(p.pattern, C.double(offset), C.double(red), C.double(green), C.double(blue))
}

// AddColorStopRGBA adds an rgba color stop to this pattern.
func (p *Pattern) AddColorStopRGBA(offset, red, green, blue, alpha float64) {
	C.cairo_pattern_add_color_stop_rgba(p.pattern, C.double(offset), C.double(red), C.double(green), C.double(blue), C.double(alpha))
}

// GetColorStopCount returns the number of color stops on this pattern.
func (p *Pattern) GetColorStopCount() int {
	var count C.int
	C.cairo_pattern_get_color_stop_count(p.pattern, &count)
	return int(count)
}

// GetColorStopRGBA returns the value for the color stop at the given index.
func (p *Pattern) GetColorStopRGBA(index int) (offset, red, green, blue, alpha float64) {
	var o C.double
	var r C.double
	var g C.double
	var b C.double
	var a C.double
	C.cairo_pattern_get_color_stop_rgba(p.pattern, C.int(index), &o, &r, &g, &b, &a)
	return float64(o), float64(r), float64(g), float64(b), float64(a)
}

// GetLinearPoints returns the values of the points defining a linear gradient.
func (p *Pattern) GetLinearPoints() (x0, y0, x1, y1 float64) {
	var cx0 C.double
	var cy0 C.double
	var cx1 C.double
	var cy1 C.double
	C.cairo_pattern_get_linear_points(p.pattern, &cx0, &cy0, &cx1, &cy1)
	return float64(cx0), float64(cy0), float64(cx1), float64(cy1)
}

// GetRadialCircles returns the values of the circles defining a radial gradient.
func (p *Pattern) GetRadialCircles() (x0, y0, r0, x1, y1, r1 float64) {
	var cx0 C.double
	var cy0 C.double
	var cr0 C.double
	var cx1 C.double
	var cy1 C.double
	var cr1 C.double
	C.cairo_pattern_get_radial_circles(p.pattern, &cx0, &cy0, &cr0, &cx1, &cy1, &cr1)
	return float64(cx0), float64(cy0), float64(cr0), float64(cx1), float64(cy1), float64(cr1)
}

////////////////////////////
// solid methods

// GetRGBA returns the rgba values for a solid pattern.
func (p *Pattern) GetRGBA() (red, green, blue, alpha float64) {
	var r C.double
	var g C.double
	var b C.double
	var a C.double
	C.cairo_pattern_get_rgba(p.pattern, &r, &g, &b, &a)
	return float64(r), float64(g), float64(b), float64(a)
}

////////////////////////////
// matrix methods

// SetMatrix transforms the pattern according to the matrix passed.
func (p *Pattern) SetMatrix(matrix *Matrix) {
	C.cairo_pattern_set_matrix(p.pattern, matrix.Native())
}

// GetMatrix returns the current matrix in use on the pattern.
func (p *Pattern) GetMatrix() *Matrix {
	var matrix C.cairo_matrix_t
	C.cairo_pattern_get_matrix(p.pattern, &matrix)
	return &Matrix{
		float64(matrix.xx),
		float64(matrix.yx),
		float64(matrix.xy),
		float64(matrix.yy),
		float64(matrix.x0),
		float64(matrix.y0),
	}
}

///////////////////////////
// mesh methods

// CreateMesh creates a mesh pattern
func CreateMesh() *Pattern {
	p := C.cairo_pattern_create_mesh()
	return &Pattern{p}
}

// BeginPatch starts a patch definition.
func (p *Pattern) BeginPatch() {
	C.cairo_mesh_pattern_begin_patch(p.pattern)
}

// EndPatch completes a patch definition.
func (p *Pattern) EndPatch() {
	C.cairo_mesh_pattern_end_patch(p.pattern)
}

// MoveTo moves to an x, y point.
func (p *Pattern) MoveTo(x, y float64) {
	C.cairo_mesh_pattern_move_to(p.pattern, C.double(x), C.double(y))
}

// LineTo draws a line to an x, y point.
func (p *Pattern) LineTo(x, y float64) {
	C.cairo_mesh_pattern_line_to(p.pattern, C.double(x), C.double(y))
}

// CurveTo draws a bezier curve to a point through two control points.
func (p *Pattern) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	C.cairo_mesh_pattern_curve_to(p.pattern, C.double(x1), C.double(y1), C.double(x2), C.double(y2), C.double(x3), C.double(y3))
}

// SetControlPoint sets the x, y position of a given control point.
func (p *Pattern) SetControlPoint(pointNum uint, x, y float64) {
	C.cairo_mesh_pattern_set_control_point(p.pattern, C.uint(pointNum), C.double(x), C.double(y))
}

// SetCornerColorRGB sets the RGB color for a given corner.
func (p *Pattern) SetCornerColorRGB(cornerNum uint, r, g, b float64) {
	C.cairo_mesh_pattern_set_corner_color_rgb(p.pattern, C.uint(cornerNum), C.double(r), C.double(g), C.double(b))
}

// SetCornerColorRGBA sets the RGBA color for a given corner.
func (p *Pattern) SetCornerColorRGBA(cornerNum uint, r, g, b, a float64) {
	C.cairo_mesh_pattern_set_corner_color_rgba(p.pattern, C.uint(cornerNum), C.double(r), C.double(g), C.double(b), C.double(a))
}

// GetPatchCount returns the number of patches defined for this pattern.
func (p *Pattern) GetPatchCount() uint {
	var count C.uint
	C.cairo_mesh_pattern_get_patch_count(p.pattern, &count)
	return uint(count)
}

// TODO: need to implement cairo_path_t
// func (p *Pattern) GetPath(patchNum uint) uint {
// 	return C.cairo_mesh_pattern_get_path(p.pattern, C.uint(patchNum))
// }

// GetControlPoint returns the control point for a given patch corner.
func (p *Pattern) GetControlPoint(patchNum, pointNum uint) (float64, float64) {
	var x C.double
	var y C.double
	C.cairo_mesh_pattern_get_control_point(p.pattern, C.uint(patchNum), C.uint(pointNum), &x, &y)
	return float64(x), float64(y)
}

// GetCornerColorRGBA returns the RGBA color for a given patch corner.
func (p *Pattern) GetCornerColorRGBA(patchNum, pointNum uint) (float64, float64, float64, float64) {
	var r C.double
	var g C.double
	var b C.double
	var a C.double
	C.cairo_mesh_pattern_get_corner_color_rgba(p.pattern, C.uint(patchNum), C.uint(pointNum), &r, &g, &b, &a)
	return float64(r), float64(g), float64(b), float64(a)
}

// TODO:
// cairo_pattern_t *	cairo_pattern_create_for_surface ()
// cairo_status_t	cairo_pattern_get_surface ()
// cairo_pattern_t *	cairo_pattern_reference ()
// void	cairo_pattern_destroy ()
// cairo_status_t	cairo_pattern_status ()
// void	cairo_pattern_set_extend ()
// cairo_extend_t	cairo_pattern_get_extend ()
// void	cairo_pattern_set_filter ()
// cairo_filter_t	cairo_pattern_get_filter ()
// cairo_pattern_type_t	cairo_pattern_get_type ()
// unsigned int	cairo_pattern_get_reference_count ()
// cairo_status_t	cairo_pattern_set_user_data ()
// void *	cairo_pattern_get_user_data ()
