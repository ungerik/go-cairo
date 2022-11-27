// Package cairo wraps the c cairographics library.
package cairo

// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

import "unsafe"

// Context represents a cairo context
type Context struct {
	context *C.cairo_t
}

// NewContext creates a new cairo context.
func NewContext(surface *Surface) *Context {
	return &Context{
		C.cairo_create(surface.surface),
	}
}

// Rectangle rectangle struct
type Rectangle struct {
	X, Y          float64
	Width, Height float64
}

// Native returns a pointer to the native cairo context
func (c Context) Native() uintptr {
	return uintptr(unsafe.Pointer(c.context))
}

// GetCurrentPoint gets the current drawing point.
func (c *Context) GetCurrentPoint() (float64, float64) {
	if !c.HasCurrentPoint() {
		return 0, 0
	}
	x := C.double(0)
	y := C.double(0)
	C.cairo_get_current_point(c.context, &x, &y)

	// if c.GetStatus() != StatusSuccess {
	// May not need to panic here. Per cairo spec, if status is error, return 0, 0, which this will do.
	// panic("cairo.Context.GetCurrentPoint() unable to get current point.")
	// }
	return float64(x), float64(y)
}

// HasCurrentPoint returns whether or not there is a current drawing point.
func (c *Context) HasCurrentPoint() bool {
	return C.cairo_has_current_point(c.context) != 0
}

// Save saves the current state of the context.
func (c *Context) Save() {
	C.cairo_save(c.context)
}

// Restore restores the the last saved state of the context.
func (c *Context) Restore() {
	C.cairo_restore(c.context)
}

// PushGroup temporarily redirects drawing to an intermediate context known as a group.
func (c *Context) PushGroup() {
	C.cairo_push_group(c.context)
}

// PushGroupWithContent temporarily redirects drawing to an intermediate context known as a group, with content.
func (c *Context) PushGroupWithContent(content Content) {
	C.cairo_push_group_with_content(c.context, C.cairo_content_t(content))
}

// PopGroup terminates the redirection begun by a call to cairo_push_group() or cairo_push_group_with_content() and returns a new pattern containing the results of all drawing operations performed to the group.
func (c *Context) PopGroup() (pattern *Pattern) {
	return &Pattern{C.cairo_pop_group(c.context)}
}

// PopGroupToSource terminates the redirection begun by a call to cairo_push_group() or cairo_push_group_with_content() and installs the resulting pattern as the source pattern in the given cairo context.
func (c *Context) PopGroupToSource() {
	C.cairo_pop_group_to_source(c.context)
}

// SetOperator sets the compositing operator to be used for all drawing operations.
func (c *Context) SetOperator(operator Operator) {
	C.cairo_set_operator(c.context, C.cairo_operator_t(operator))
}

// SetSource sets the pattern to draw with.
func (c *Context) SetSource(pattern *Pattern) {
	C.cairo_set_source(c.context, pattern.pattern)
}

// GetSource gets the current pattern being used.
func (c *Context) GetSource() *Pattern {
	return &Pattern{C.cairo_get_source(c.context)}
}

// SetSourceRGB sets the r, g, b values to draw with.
func (c *Context) SetSourceRGB(red, green, blue float64) {
	C.cairo_set_source_rgb(c.context, C.double(red), C.double(green), C.double(blue))
}

// SetSourceRGBA sets the r, g, b, a values to draw with.
func (c *Context) SetSourceRGBA(red, green, blue, alpha float64) {
	C.cairo_set_source_rgba(c.context, C.double(red), C.double(green), C.double(blue), C.double(alpha))
}

// SetSourceSurface is a convenience function for creating a pattern from surface and setting it as the source in cr with cairo_set_source().
func (c *Context) SetSourceSurface(surface *Surface, x, y float64) {
	C.cairo_set_source_surface(c.context, surface.surface, C.double(x), C.double(y))
}

// SetTolerance sets the tolerance used when converting paths into trapezoids.
func (c *Context) SetTolerance(tolerance float64) {
	C.cairo_set_tolerance(c.context, C.double(tolerance))
}

// SetAntialias sets the antialias value to use.
func (c *Context) SetAntialias(antialias Antialias) {
	C.cairo_set_antialias(c.context, C.cairo_antialias_t(antialias))
}

// SetFillRule sets the current fill rule within the cairo context.
func (c *Context) SetFillRule(fillRule FillRule) {
	C.cairo_set_fill_rule(c.context, C.cairo_fill_rule_t(fillRule))
}

// SetLineWidth sets the pixel width that will be used when drawing lines.
func (c *Context) SetLineWidth(width float64) {
	C.cairo_set_line_width(c.context, C.double(width))
}

// GetLineWidth sets the pixel width that will be used when drawing lines.
func (c *Context) GetLineWidth() float64 {
	w := C.cairo_get_line_width(c.context)
	return float64(w)
}

// SetLineCap sets the form of line cap used when drawing lines.
func (c *Context) SetLineCap(lineCap LineCap) {
	C.cairo_set_line_cap(c.context, C.cairo_line_cap_t(lineCap))
}

// SetLineJoin sets the type of join to use where two line segments connect.
func (c *Context) SetLineJoin(lineJoin LineJoin) {
	C.cairo_set_line_join(c.context, C.cairo_line_join_t(lineJoin))
}

// SetDash sets the dash pattern to be used when drawing lines.
func (c *Context) SetDash(dashes []float64, numDashes int, offset float64) {
	dashesp := (*C.double)(&dashes[0])
	C.cairo_set_dash(c.context, dashesp, C.int(numDashes), C.double(offset))
}

// SetMiterLimit sets the sharpness of the corner in line joins.
func (c *Context) SetMiterLimit(limit float64) {
	C.cairo_set_miter_limit(c.context, C.double(limit))
}

// Translate translates the context by the specified amounts.
func (c *Context) Translate(tx, ty float64) {
	C.cairo_translate(c.context, C.double(tx), C.double(ty))
}

// Scale scales the context by the specified amount.
func (c *Context) Scale(sx, sy float64) {
	C.cairo_scale(c.context, C.double(sx), C.double(sy))
}

// Rotate rotates the context by the specified amount.
func (c *Context) Rotate(angle float64) {
	C.cairo_rotate(c.context, C.double(angle))
}

// Transform transforms the context with the specified matrix.
func (c *Context) Transform(matrix Matrix) {
	C.cairo_transform(c.context, matrix.Native())
}

// SetMatrix resets the context transform to the specified matrix
func (c *Context) SetMatrix(matrix Matrix) {
	C.cairo_set_matrix(c.context, matrix.Native())
}

// IdentityMatrix sets the transformation matrix for the context to an identity matrix.
func (c *Context) IdentityMatrix() {
	C.cairo_identity_matrix(c.context)
}

// UserToDevice tbd
func (c *Context) UserToDevice(x, y float64) (float64, float64) {
	C.cairo_user_to_device(c.context, (*C.double)(&x), (*C.double)(&y))
	return x, y
}

// UserToDeviceDistance tbd
func (c *Context) UserToDeviceDistance(dx, dy float64) (float64, float64) {
	C.cairo_user_to_device_distance(c.context, (*C.double)(&dx), (*C.double)(&dy))
	return dx, dy
}

// DeviceToUser tbd
func (c *Context) DeviceToUser(x, y float64) (float64, float64) {
	C.cairo_device_to_user(c.context, (*C.double)(&x), (*C.double)(&y))
	return x, y
}

// DeviceToUserDistance tbd
func (c *Context) DeviceToUserDistance(x, y float64) (float64, float64) {
	C.cairo_device_to_user_distance(c.context, (*C.double)(&x), (*C.double)(&y))
	return x, y
}

// path creation methods

// NewPath begins a new drawing path.
func (c *Context) NewPath() {
	C.cairo_new_path(c.context)
}

// MoveTo moves to the specified point.
func (c *Context) MoveTo(x, y float64) {
	C.cairo_move_to(c.context, C.double(x), C.double(y))
}

// NewSubPath creates a new sub drawing path.
func (c *Context) NewSubPath() {
	C.cairo_new_sub_path(c.context)
}

// LineTo draws a line to the specified point.
func (c *Context) LineTo(x, y float64) {
	C.cairo_line_to(c.context, C.double(x), C.double(y))
}

// CurveTo draws a Bezier curve through the specified points.
func (c *Context) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	C.cairo_curve_to(c.context,
		C.double(x1), C.double(y1),
		C.double(x2), C.double(y2),
		C.double(x3), C.double(y3))
}

// Arc draws and arc with the specified parameters.
func (c *Context) Arc(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc(c.context,
		C.double(xc), C.double(yc),
		C.double(radius),
		C.double(angle1), C.double(angle2))
}

// ArcNegative draws a negative arc to the specified parameters.
func (c *Context) ArcNegative(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc_negative(c.context,
		C.double(xc), C.double(yc),
		C.double(radius),
		C.double(angle1), C.double(angle2))
}

// RelMoveTo tbd
func (c *Context) RelMoveTo(dx, dy float64) {
	C.cairo_rel_move_to(c.context, C.double(dx), C.double(dy))
}

// RelLineTo tbd
func (c *Context) RelLineTo(dx, dy float64) {
	C.cairo_rel_line_to(c.context, C.double(dx), C.double(dy))
}

// RelCurveTo tbd
func (c *Context) RelCurveTo(dx1, dy1, dx2, dy2, dx3, dy3 float64) {
	C.cairo_rel_curve_to(c.context,
		C.double(dx1), C.double(dy1),
		C.double(dx2), C.double(dy2),
		C.double(dx3), C.double(dy3))
}

// Rectangle tbd
func (c *Context) Rectangle(x, y, width, height float64) {
	C.cairo_rectangle(c.context,
		C.double(x), C.double(y),
		C.double(width), C.double(height))
}

// ClosePath tbd
func (c *Context) ClosePath() {
	C.cairo_close_path(c.context)
}

// PathExtents tbd
func (c *Context) PathExtents() (left, top, right, bottom float64) {
	C.cairo_path_extents(c.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}

///////////////////////////////////////////////////////////////////////////////
// Painting methods

// Paint tbd
func (c *Context) Paint() {
	C.cairo_paint(c.context)
}

// PaintWithAlpha tbd
func (c *Context) PaintWithAlpha(alpha float64) {
	C.cairo_paint_with_alpha(c.context, C.double(alpha))
}

// Mask tbd
func (c *Context) Mask(pattern Pattern) {
	C.cairo_mask(c.context, pattern.pattern)
}

// MaskSurface tbd
func (c *Context) MaskSurface(surface *Surface, surfaceX, surfaceY float64) {
	C.cairo_mask_surface(c.context, surface.surface, C.double(surfaceX), C.double(surfaceY))
}

// Stroke tbd
func (c *Context) Stroke() {
	C.cairo_stroke(c.context)
}

// StrokePreserve tbd
func (c *Context) StrokePreserve() {
	C.cairo_stroke_preserve(c.context)
}

// Fill tbd
func (c *Context) Fill() {
	C.cairo_fill(c.context)
}

// FillPreserve tbd
func (c *Context) FillPreserve() {
	C.cairo_fill_preserve(c.context)
}

// CopyPage tbd
func (c *Context) CopyPage() {
	C.cairo_copy_page(c.context)
}

// ShowPage tbd
func (c *Context) ShowPage() {
	C.cairo_show_page(c.context)
}

///////////////////////////////////////////////////////////////////////////////
// Insideness testing

// InStroke tbd
func (c *Context) InStroke(x, y float64) bool {
	return C.cairo_in_stroke(c.context, C.double(x), C.double(y)) != 0
}

// InFill tbd
func (c *Context) InFill(x, y float64) bool {
	return C.cairo_in_fill(c.context, C.double(x), C.double(y)) != 0
}

///////////////////////////////////////////////////////////////////////////////
// Rectangular extents

// StrokeExtents tbd
func (c *Context) StrokeExtents() (left, top, right, bottom float64) {
	C.cairo_stroke_extents(c.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}

// FillExtents tbd
func (c *Context) FillExtents() (left, top, right, bottom float64) {
	C.cairo_fill_extents(c.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}

///////////////////////////////////////////////////////////////////////////////
// Clipping methods

// ResetClip tbd
func (c *Context) ResetClip() {
	C.cairo_reset_clip(c.context)
}

// Clip tbd
func (c *Context) Clip() {
	C.cairo_clip(c.context)
}

// ClipPreserve tbd
func (c *Context) ClipPreserve() {
	C.cairo_clip_preserve(c.context)
}

// ClipExtents tbd
func (c *Context) ClipExtents() (left, top, right, bottom float64) {
	C.cairo_clip_extents(c.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}

// ClipRectangleList tbd
func (c *Context) ClipRectangleList() ([]Rectangle, Status) {
	list := C.cairo_copy_clip_rectangle_list(c.context)
	defer C.cairo_rectangle_list_destroy(list)
	rects := make([]Rectangle, int(list.num_rectangles))
	C.memcpy(unsafe.Pointer(&rects[0]), unsafe.Pointer(list.rectangles), C.size_t(list.num_rectangles*8))
	return rects, Status(list.status)
}

///////////////////////////////////////////////////////////////////////////////
// Font/Text methods

// SelectFontFace tbd
func (c *Context) SelectFontFace(name string, fontSlant, fontWeight int) {
	str := C.CString(name)
	C.cairo_select_font_face(c.context, str, C.cairo_font_slant_t(fontSlant), C.cairo_font_weight_t(fontWeight))
	C.free(unsafe.Pointer(str))
}

// SetFontSize tbd
func (c *Context) SetFontSize(size float64) {
	C.cairo_set_font_size(c.context, C.double(size))
}

// SetFontMatrix tbd
func (c *Context) SetFontMatrix(matrix Matrix) {
	C.cairo_set_font_matrix(c.context, matrix.Native())
}

// SetFontOptions tbd
func (c *Context) SetFontOptions(fontOptions *FontOptions) {
	panic("not implemented") // todo
}

// GetFontOptions tbd
func (c *Context) GetFontOptions() *FontOptions {
	panic("not implemented") // todo
}

// SetFontFace tbd
func (c *Context) SetFontFace(fontFace *FontFace) {
	panic("not implemented") // todo
}

// GetFontFace tbd
func (c *Context) GetFontFace() *FontFace {
	panic("not implemented") // todo
}

// SetScaledFont tbd
func (c *Context) SetScaledFont(scaledFont *ScaledFont) {
	panic("not implemented") // todo
}

// GetScaledFont tbd
func (c *Context) GetScaledFont() *ScaledFont {
	panic("not implemented") // todo
}

// ShowText tbd
func (c *Context) ShowText(text string) {
	cs := C.CString(text)
	C.cairo_show_text(c.context, cs)
	C.free(unsafe.Pointer(cs))
}

// ShowGlyphs tbd
func (c *Context) ShowGlyphs(glyphs []Glyph) {
	panic("not implemented") // todo
}

// ShowTextGlyphs tbd
func (c *Context) ShowTextGlyphs(text string, glyphs []Glyph, clusters []TextCluster, flags TextClusterFlag) {
}

// TextPath tbd
func (c *Context) TextPath(text string) {
	cs := C.CString(text)
	C.cairo_text_path(c.context, cs)
	C.free(unsafe.Pointer(cs))
}

// GlyphPath tbd
func (c *Context) GlyphPath(glyphs []Glyph) {
	panic("not implemented") // todo
}

// TextExtents tbd
func (c *Context) TextExtents(text string) *TextExtents {
	cte := C.cairo_text_extents_t{}
	cs := C.CString(text)
	C.cairo_text_extents(c.context, cs, &cte)
	C.free(unsafe.Pointer(cs))
	te := &TextExtents{
		Xbearing: float64(cte.x_bearing),
		Ybearing: float64(cte.y_bearing),
		Width:    float64(cte.width),
		Height:   float64(cte.height),
		Xadvance: float64(cte.x_advance),
		Yadvance: float64(cte.y_advance),
	}
	return te
}

// GlyphExtents tbd
func (c *Context) GlyphExtents(glyphs []Glyph) *TextExtents {
	panic("not implemented") // todo
	//C.cairo_text_extents
}

// FontExtents tbd
func (c *Context) FontExtents() *FontExtents {
	panic("not implemented") // todo
	//C.cairo_text_extents
}

// GetStatus tbd
func (c *Context) GetStatus() Status {
	return Status(C.cairo_status(c.context))
}

// Destroy tbd
func (c *Context) Destroy() {
	C.cairo_destroy(c.context)
}
