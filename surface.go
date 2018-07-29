// +build !goci

package cairo

// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"image"
	"image/draw"
	"unsafe"

	"github.com/bit101/go-cairo/extimage"
)

// Surface Golang struct to hold both a cairo surface and a cairo context
type Surface struct {
	surface *C.cairo_surface_t
	context *C.cairo_t
}

// NewSurface creates a new Surface struct.
func NewSurface(format Format, width, height int) *Surface {
	s := C.cairo_image_surface_create(C.cairo_format_t(format), C.int(width), C.int(height))
	return &Surface{surface: s, context: C.cairo_create(s)}
}

// NewSurfaceFromC creates a new surface from C data types.
// This is useful, if you already obtained a surface by
// using a C library, for example an XCB surface.
func NewSurfaceFromC(s *C.cairo_surface_t, c *C.cairo_t) *Surface {
	return &Surface{surface: s, context: c}
}

// NewSurfaceFromPNG creates a new Surface struct from a png file.
func NewSurfaceFromPNG(filename string) (*Surface, Status) {

	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))

	surfaceNative := C.cairo_image_surface_create_from_png(cstr)
	status := Status(C.cairo_surface_status(surfaceNative))
	if status != StatusSuccess {
		return nil, status
	}

	contextNative := C.cairo_create(surfaceNative)
	status = Status(C.cairo_status(contextNative))
	if status != StatusSuccess {
		return nil, status
	}

	surface := &Surface{
		surface: surfaceNative,
		context: contextNative,
	}

	return surface, StatusSuccess
}

// Native returns a pointer to the native cairo surface and context.
func (s *Surface) Native() (surface, context uintptr) {
	surface = uintptr(unsafe.Pointer(s.surface))
	context = uintptr(unsafe.Pointer(s.context))

	return
}

// NewSurfaceFromImage created a new Surface struct from an Image object.
func NewSurfaceFromImage(img image.Image) *Surface {
	var format Format
	switch img.(type) {
	case *image.Alpha, *image.Alpha16:
		format = FormatA8
	case *extimage.BGRN, *image.Gray, *image.Gray16, *image.YCbCr:
		format = FormatRGB24
	default:
		format = FormatARGB32
	}
	surface := NewSurface(format, img.Bounds().Dx(), img.Bounds().Dy())
	surface.SetImage(img)
	return surface
}

// NewPDFSurface creates a new Surface struct used for saving as a PDF.
func NewPDFSurface(filename string, widthInPoints, heightInPoints float64, version PDFVersion) *Surface {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	s := C.cairo_pdf_surface_create(cs, C.double(widthInPoints), C.double(heightInPoints))
	C.cairo_pdf_surface_restrict_to_version(s, C.cairo_pdf_version_t(version))
	return &Surface{surface: s, context: C.cairo_create(s)}
}

// NewPSSurface creates a new Surface sturct used in saving as a postscript file.
func NewPSSurface(filename string, widthInPoints, heightInPoints float64, level PSLevel) *Surface {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	s := C.cairo_ps_surface_create(cs, C.double(widthInPoints), C.double(heightInPoints))
	C.cairo_ps_surface_restrict_to_level(s, C.cairo_ps_level_t(level))
	return &Surface{surface: s, context: C.cairo_create(s)}
}

// NewSVGSurface creates a new Surface struct used in saving as an SVG file.
func NewSVGSurface(filename string, widthInPoints, heightInPoints float64, version SVGVersion) *Surface {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	s := C.cairo_svg_surface_create(cs, C.double(widthInPoints), C.double(heightInPoints))
	C.cairo_svg_surface_restrict_to_version(s, C.cairo_svg_version_t(version))
	return &Surface{surface: s, context: C.cairo_create(s)}
}

// GetCurrentPoint gets the current drawing point.
func (s *Surface) GetCurrentPoint() (float64, float64) {
	if !s.HasCurrentPoint() {
		return 0, 0
	}
	x := C.double(0)
	y := C.double(0)
	C.cairo_get_current_point(s.context, &x, &y)
	if s.GetStatus() != StatusSuccess {
		// May not need to panic here. Per cairo spec, if status is error, return 0, 0, which this will do.
		panic("cairo.Surface.GetCurrentPoint() unable to get current point.")
	}
	return float64(x), float64(y)
}

// HasCurrentPoint returns whether or not there is a current drawing point.
func (s *Surface) HasCurrentPoint() bool {
	return C.cairo_has_current_point(s.context) != 0
}

// Save saves the current state of the context.
func (s *Surface) Save() {
	C.cairo_save(s.context)
}

// Restore restores the the last saved state of the context.
func (s *Surface) Restore() {
	C.cairo_restore(s.context)
}

// PushGroup ...
func (s *Surface) PushGroup() {
	C.cairo_push_group(s.context)
}

// PushGroupWithContent ...
func (s *Surface) PushGroupWithContent(content Content) {
	C.cairo_push_group_with_content(s.context, C.cairo_content_t(content))
}

// PopGroup ...
func (s *Surface) PopGroup() (pattern *Pattern) {
	return &Pattern{C.cairo_pop_group(s.context)}
}

// PopGroupToSource ...
func (s *Surface) PopGroupToSource() {
	C.cairo_pop_group_to_source(s.context)
}

// SetOperator ...
func (s *Surface) SetOperator(operator Operator) {
	C.cairo_set_operator(s.context, C.cairo_operator_t(operator))
}

// SetSource sets the pattern to draw with.
func (s *Surface) SetSource(pattern *Pattern) {
	C.cairo_set_source(s.context, pattern.pattern)
}

// SetSourceRGB sets the r, g, b values to draw with.
func (s *Surface) SetSourceRGB(red, green, blue float64) {
	C.cairo_set_source_rgb(s.context, C.double(red), C.double(green), C.double(blue))
}

// SetSourceRGBA sets the r, g, b, a values to draw with.
func (s *Surface) SetSourceRGBA(red, green, blue, alpha float64) {
	C.cairo_set_source_rgba(s.context, C.double(red), C.double(green), C.double(blue), C.double(alpha))
}

// SetSourceSurface ...
func (s *Surface) SetSourceSurface(surface *Surface, x, y float64) {
	C.cairo_set_source_surface(s.context, surface.surface, C.double(x), C.double(y))
}

// SetTolerance ...
func (s *Surface) SetTolerance(tolerance float64) {
	C.cairo_set_tolerance(s.context, C.double(tolerance))
}

// SetAntialias sets the antialias value to use.
func (s *Surface) SetAntialias(antialias Antialias) {
	C.cairo_set_antialias(s.context, C.cairo_antialias_t(antialias))
}

// SetFillRule ...
func (s *Surface) SetFillRule(fillRule FillRule) {
	C.cairo_set_fill_rule(s.context, C.cairo_fill_rule_t(fillRule))
}

// SetLineWidth sets the pixel width that will be used when drawing lines.
func (s *Surface) SetLineWidth(width float64) {
	C.cairo_set_line_width(s.context, C.double(width))
}

// SetLineCap sets the form of line cap used when drawing lines.
func (s *Surface) SetLineCap(lineCap LineCap) {
	C.cairo_set_line_cap(s.context, C.cairo_line_cap_t(lineCap))
}

// SetLineJoin sets the type of join to use where two line segments connect.
func (s *Surface) SetLineJoin(lineJoin LineJoin) {
	C.cairo_set_line_join(s.context, C.cairo_line_join_t(lineJoin))
}

// SetDash sets the dash pattern to be used when drawing lines.
func (s *Surface) SetDash(dashes []float64, numDashes int, offset float64) {
	dashesp := (*C.double)(&dashes[0])
	C.cairo_set_dash(s.context, dashesp, C.int(numDashes), C.double(offset))
}

// SetMiterLimit sets the sharpness of the corner in line joins.
func (s *Surface) SetMiterLimit(limit float64) {
	C.cairo_set_miter_limit(s.context, C.double(limit))
}

// Translate translates the surface by the specified amounts.
func (s *Surface) Translate(tx, ty float64) {
	C.cairo_translate(s.context, C.double(tx), C.double(ty))
}

// Scale scales the surface by the specified amount.
func (s *Surface) Scale(sx, sy float64) {
	C.cairo_scale(s.context, C.double(sx), C.double(sy))
}

// Rotate rotates the surface by the specified amount.
func (s *Surface) Rotate(angle float64) {
	C.cairo_rotate(s.context, C.double(angle))
}

// Transform transforms the surface with the specified matrix.
func (s *Surface) Transform(matrix Matrix) {
	C.cairo_transform(s.context, matrix.Native())
}

// SetMatrix resets the surface transform to the specified matrix
func (s *Surface) SetMatrix(matrix Matrix) {
	C.cairo_set_matrix(s.context, matrix.Native())
}

// IdentityMatrix ...
func (s *Surface) IdentityMatrix() {
	C.cairo_identity_matrix(s.context)
}

// UserToDevice ...
func (s *Surface) UserToDevice(x, y float64) (float64, float64) {
	C.cairo_user_to_device(s.context, (*C.double)(&x), (*C.double)(&y))
	return x, y
}

// UserToDeviceDistance ...
func (s *Surface) UserToDeviceDistance(dx, dy float64) (float64, float64) {
	C.cairo_user_to_device_distance(s.context, (*C.double)(&dx), (*C.double)(&dy))
	return dx, dy
}

// DeviceToUser ...
func (s *Surface) DeviceToUser(x, y float64) (float64, float64) {
	C.cairo_device_to_user(s.context, (*C.double)(&x), (*C.double)(&y))
	return x, y
}

// DeviceToUserDistance ...
func (s *Surface) DeviceToUserDistance(x, y float64) (float64, float64) {
	C.cairo_device_to_user_distance(s.context, (*C.double)(&x), (*C.double)(&y))
	return x, y
}

// path creation methods

// NewPath begins a new drawing path.
func (s *Surface) NewPath() {
	C.cairo_new_path(s.context)
}

// MoveTo moves to the specified point.
func (s *Surface) MoveTo(x, y float64) {
	C.cairo_move_to(s.context, C.double(x), C.double(y))
}

// NewSubPath creates a new sub drawing path.
func (s *Surface) NewSubPath() {
	C.cairo_new_sub_path(s.context)
}

// LineTo draws a line to the specified point.
func (s *Surface) LineTo(x, y float64) {
	C.cairo_line_to(s.context, C.double(x), C.double(y))
}

// CurveTo draws a Bezier curve through the specified points.
func (s *Surface) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	C.cairo_curve_to(s.context,
		C.double(x1), C.double(y1),
		C.double(x2), C.double(y2),
		C.double(x3), C.double(y3))
}

// Arc draws and arc with the specified parameters.
func (s *Surface) Arc(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc(s.context,
		C.double(xc), C.double(yc),
		C.double(radius),
		C.double(angle1), C.double(angle2))
}

// ArcNegative draws a negative arc to the specified parameters.
func (s *Surface) ArcNegative(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc_negative(s.context,
		C.double(xc), C.double(yc),
		C.double(radius),
		C.double(angle1), C.double(angle2))
}

// RelMoveTo ...
func (s *Surface) RelMoveTo(dx, dy float64) {
	C.cairo_rel_move_to(s.context, C.double(dx), C.double(dy))
}

// RelLineTo ...
func (s *Surface) RelLineTo(dx, dy float64) {
	C.cairo_rel_line_to(s.context, C.double(dx), C.double(dy))
}

// RelCurveTo ...
func (s *Surface) RelCurveTo(dx1, dy1, dx2, dy2, dx3, dy3 float64) {
	C.cairo_rel_curve_to(s.context,
		C.double(dx1), C.double(dy1),
		C.double(dx2), C.double(dy2),
		C.double(dx3), C.double(dy3))
}

// Rectangle ...
func (s *Surface) Rectangle(x, y, width, height float64) {
	C.cairo_rectangle(s.context,
		C.double(x), C.double(y),
		C.double(width), C.double(height))
}

// ClosePath ...
func (s *Surface) ClosePath() {
	C.cairo_close_path(s.context)
}

// PathExtents ...
func (s *Surface) PathExtents() (left, top, right, bottom float64) {
	C.cairo_path_extents(s.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}

///////////////////////////////////////////////////////////////////////////////
// Painting methods

// Paint ...
func (s *Surface) Paint() {
	C.cairo_paint(s.context)
}

// PaintWithAlpha ...
func (s *Surface) PaintWithAlpha(alpha float64) {
	C.cairo_paint_with_alpha(s.context, C.double(alpha))
}

// Mask ...
func (s *Surface) Mask(pattern Pattern) {
	C.cairo_mask(s.context, pattern.pattern)
}

// MaskSurface ...
func (s *Surface) MaskSurface(surface *Surface, surfaceX, surfaceY float64) {
	C.cairo_mask_surface(s.context, surface.surface, C.double(surfaceX), C.double(surfaceY))
}

// Stroke ...
func (s *Surface) Stroke() {
	C.cairo_stroke(s.context)
}

// StrokePreserve ...
func (s *Surface) StrokePreserve() {
	C.cairo_stroke_preserve(s.context)
}

// Fill ...
func (s *Surface) Fill() {
	C.cairo_fill(s.context)
}

// FillPreserve ...
func (s *Surface) FillPreserve() {
	C.cairo_fill_preserve(s.context)
}

// CopyPage ...
func (s *Surface) CopyPage() {
	C.cairo_copy_page(s.context)
}

// ShowPage ...
func (s *Surface) ShowPage() {
	C.cairo_show_page(s.context)
}

///////////////////////////////////////////////////////////////////////////////
// Insideness testing

// InStroke ...
func (s *Surface) InStroke(x, y float64) bool {
	return C.cairo_in_stroke(s.context, C.double(x), C.double(y)) != 0
}

// InFill ...
func (s *Surface) InFill(x, y float64) bool {
	return C.cairo_in_fill(s.context, C.double(x), C.double(y)) != 0
}

///////////////////////////////////////////////////////////////////////////////
// Rectangular extents

// StrokeExtents ...
func (s *Surface) StrokeExtents() (left, top, right, bottom float64) {
	C.cairo_stroke_extents(s.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}

// FillExtents ...
func (s *Surface) FillExtents() (left, top, right, bottom float64) {
	C.cairo_fill_extents(s.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}

///////////////////////////////////////////////////////////////////////////////
// Clipping methods

// ResetClip ...
func (s *Surface) ResetClip() {
	C.cairo_reset_clip(s.context)
}

// Clip ...
func (s *Surface) Clip() {
	C.cairo_clip(s.context)
}

// ClipPreserve ...
func (s *Surface) ClipPreserve() {
	C.cairo_clip_preserve(s.context)
}

// ClipExtents ...
func (s *Surface) ClipExtents() (left, top, right, bottom float64) {
	C.cairo_clip_extents(s.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}

// ClipRectangleList ...
func (s *Surface) ClipRectangleList() ([]Rectangle, Status) {
	list := C.cairo_copy_clip_rectangle_list(s.context)
	defer C.cairo_rectangle_list_destroy(list)
	rects := make([]Rectangle, int(list.num_rectangles))
	C.memcpy(unsafe.Pointer(&rects[0]), unsafe.Pointer(list.rectangles), C.size_t(list.num_rectangles*8))
	return rects, Status(list.status)
}

///////////////////////////////////////////////////////////////////////////////
// Font/Text methods

// SelectFontFace ...
func (s *Surface) SelectFontFace(name string, fontSlant, fontWeight int) {
	str := C.CString(name)
	C.cairo_select_font_face(s.context, str, C.cairo_font_slant_t(fontSlant), C.cairo_font_weight_t(fontWeight))
	C.free(unsafe.Pointer(s))
}

// SetFontSize ...
func (s *Surface) SetFontSize(size float64) {
	C.cairo_set_font_size(s.context, C.double(size))
}

// SetFontMatrix ...
func (s *Surface) SetFontMatrix(matrix Matrix) {
	C.cairo_set_font_matrix(s.context, matrix.Native())
}

// SetFontOptions ...
func (s *Surface) SetFontOptions(fontOptions *FontOptions) {
	panic("not implemented") // todo
}

// GetFontOptions ...
func (s *Surface) GetFontOptions() *FontOptions {
	panic("not implemented") // todo
}

// SetFontFace ...
func (s *Surface) SetFontFace(fontFace *FontFace) {
	panic("not implemented") // todo
}

// GetFontFace ...
func (s *Surface) GetFontFace() *FontFace {
	panic("not implemented") // todo
}

// SetScaledFont ...
func (s *Surface) SetScaledFont(scaledFont *ScaledFont) {
	panic("not implemented") // todo
}

// GetScaledFont ...
func (s *Surface) GetScaledFont() *ScaledFont {
	panic("not implemented") // todo
}

// ShowText ...
func (s *Surface) ShowText(text string) {
	cs := C.CString(text)
	C.cairo_show_text(s.context, cs)
	C.free(unsafe.Pointer(cs))
}

// ShowGlyphs ...
func (s *Surface) ShowGlyphs(glyphs []Glyph) {
	panic("not implemented") // todo
}

// ShowTextGlyphs ...
func (s *Surface) ShowTextGlyphs(text string, glyphs []Glyph, clusters []TextCluster, flags TextClusterFlag) {
}

// TextPath ...
func (s *Surface) TextPath(text string) {
	cs := C.CString(text)
	C.cairo_text_path(s.context, cs)
	C.free(unsafe.Pointer(cs))
}

// GlyphPath ...
func (s *Surface) GlyphPath(glyphs []Glyph) {
	panic("not implemented") // todo
}

// TextExtents ...
func (s *Surface) TextExtents(text string) *TextExtents {
	cte := C.cairo_text_extents_t{}
	cs := C.CString(text)
	C.cairo_text_extents(s.context, cs, &cte)
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

// GlyphExtents ...
func (s *Surface) GlyphExtents(glyphs []Glyph) *TextExtents {
	panic("not implemented") // todo
	//C.cairo_text_extents
}

// FontExtents ...
func (s *Surface) FontExtents() *FontExtents {
	panic("not implemented") // todo
	//C.cairo_text_extents
}

///////////////////////////////////////////////////////////////////////////////
// Error status queries

// Status ...
func (s *Surface) Status() Status {
	return Status(C.cairo_status(s.context))
}

///////////////////////////////////////////////////////////////////////////////
// Backend device manipulation

///////////////////////////////////////////////////////////////////////////////
// Surface manipulation

// CreateForRectangle ...
func (s *Surface) CreateForRectangle(x, y, width, height float64) *Surface {
	return &Surface{
		context: s.context,
		surface: C.cairo_surface_create_for_rectangle(s.surface,
			C.double(x), C.double(y), C.double(width), C.double(height)),
	}
}

// Finish ...
func (s *Surface) Finish() {
	C.cairo_surface_finish(s.surface)
}

// Destroy ...
func (s *Surface) Destroy() {
	C.cairo_destroy(s.context)
	C.cairo_surface_destroy(s.surface)
}

// GetDevice ...
func (s *Surface) GetDevice() *Device {
	//C.cairo_surface_get_device
	panic("not implemented") // todo
}

// GetReferenceCount ...
func (s *Surface) GetReferenceCount() int {
	return int(C.cairo_surface_get_reference_count(s.surface))
}

// GetStatus ...
func (s *Surface) GetStatus() Status {
	return Status(C.cairo_surface_status(s.surface))
}

// GetType ...
func (s *Surface) GetType() SurfaceType {
	return SurfaceType(C.cairo_surface_get_type(s.surface))
}

// GetContent ...
func (s *Surface) GetContent() Content {
	return Content(C.cairo_surface_get_content(s.surface))
}

// WriteToPNG ...
func (s *Surface) WriteToPNG(filename string) Status {

	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))

	return Status(C.cairo_surface_write_to_png(s.surface, cs))
}

// Already implemented via context split context/surface?
// ) ...
// func (s *Surface) GetFontOptions() *FontOptions {
// 	// todo
// 	// C.cairo_surface_get_font_options (cairo_surface_t      *surface,				cairo_font_options_t *options);
// 	return nil
// }

// Flush ...
func (s *Surface) Flush() {
	C.cairo_surface_flush(s.surface)
}

// MarkDirty ...
func (s *Surface) MarkDirty() {
	C.cairo_surface_mark_dirty(s.surface)
}

// MarkDirtyRectangle ...
func (s *Surface) MarkDirtyRectangle(x, y, width, height int) {
	C.cairo_surface_mark_dirty_rectangle(s.surface,
		C.int(x), C.int(y), C.int(width), C.int(height))
}

// SetDeviceOffset ...
func (s *Surface) SetDeviceOffset(x, y float64) {
	C.cairo_surface_set_device_offset(s.surface, C.double(x), C.double(y))
}

// GetDeviceOffset ...
func (s *Surface) GetDeviceOffset() (x, y float64) {
	C.cairo_surface_get_device_offset(s.surface, (*C.double)(&x), (*C.double)(&y))
	return x, y
}

// SetFallbackResolution ...
func (s *Surface) SetFallbackResolution(xPixelPerInch, yPixelPerInch float64) {
	C.cairo_surface_set_fallback_resolution(s.surface,
		C.double(xPixelPerInch), C.double(yPixelPerInch))
}

// GetFallbackResolution ...
func (s *Surface) GetFallbackResolution() (xPixelPerInch, yPixelPerInch float64) {
	C.cairo_surface_get_fallback_resolution(s.surface,
		(*C.double)(&xPixelPerInch), (*C.double)(&yPixelPerInch))
	return xPixelPerInch, yPixelPerInch
}

// Already defined for context
// ) ...
// func (s *Surface) CopyPage() {
// 	C.cairo_surface_copy_page(s.surface)
// }

// ) ...
// func (s *Surface) ShowPage() {
// 	C.cairo_surface_show_page(s.surface)
// }

// HasShowTextGlyphs ...
func (s *Surface) HasShowTextGlyphs() bool {
	return C.cairo_surface_has_show_text_glyphs(s.surface) != 0
}

// GetData returns a copy of the surfaces raw pixel data.
// This method also calls Flush.
// GetData ...
func (s *Surface) GetData() []byte {
	s.Flush()
	dataPtr := C.cairo_image_surface_get_data(s.surface)
	if dataPtr == nil {
		panic("cairo.Surface.GetData(): can't access surface pixel data")
	}
	stride := C.cairo_image_surface_get_stride(s.surface)
	height := C.cairo_image_surface_get_height(s.surface)
	return C.GoBytes(unsafe.Pointer(dataPtr), stride*height)
}

// SetData sets the surfaces raw pixel data.
// This method also calls Flush and MarkDirty.
// SetData ...
func (s *Surface) SetData(data []byte) {
	s.Flush()
	dataPtr := unsafe.Pointer(C.cairo_image_surface_get_data(s.surface))
	if dataPtr == nil {
		panic("cairo.Surface.SetData(): can't access surface pixel data")
	}
	stride := C.cairo_image_surface_get_stride(s.surface)
	height := C.cairo_image_surface_get_height(s.surface)
	if len(data) != int(stride*height) {
		panic("cairo.Surface.SetData(): invalid data size")
	}
	C.memcpy(dataPtr, unsafe.Pointer(&data[0]), C.size_t(stride*height))
	s.MarkDirty()
}

// GetFormat ...
func (s *Surface) GetFormat() Format {
	return Format(C.cairo_image_surface_get_format(s.surface))
}

// GetWidth ...
func (s *Surface) GetWidth() int {
	return int(C.cairo_image_surface_get_width(s.surface))
}

// GetHeight ...
func (s *Surface) GetHeight() int {
	return int(C.cairo_image_surface_get_height(s.surface))
}

// GetStride ...
func (s *Surface) GetStride() int {
	return int(C.cairo_image_surface_get_stride(s.surface))
}

///////////////////////////////////////////////////////////////////////////////
// image.Image methods

// GetImage returns an image based on this surface's data.
func (s *Surface) GetImage() image.Image {
	width := s.GetWidth()
	height := s.GetHeight()
	stride := s.GetStride()
	data := s.GetData()

	switch s.GetFormat() {
	case FormatARGB32:
		return &extimage.BGRA{
			Pix:    data,
			Stride: stride,
			Rect:   image.Rect(0, 0, width, height),
		}

	case FormatRGB24:
		return &extimage.BGRN{
			Pix:    data,
			Stride: stride,
			Rect:   image.Rect(0, 0, width, height),
		}

	case FormatA8:
		return &image.Alpha{
			Pix:    data,
			Stride: stride,
			Rect:   image.Rect(0, 0, width, height),
		}

	case FormatA1:
		panic("Unsuppored surface format cairo.FormatA1")

	case FormatRGB16565:
		panic("Unsuppored surface format cairo.FormatRGB16565")

	case FormatRGB30:
		panic("Unsuppored surface format cairo.FormatRGB30")

	case FormatInvalid:
		panic("Invalid surface format")
	}
	panic("Unknown surface format")
}

// SetImage set the data from an image.Image with identical size.
func (s *Surface) SetImage(img image.Image) {
	width := s.GetWidth()
	height := s.GetHeight()
	stride := s.GetStride()

	switch s.GetFormat() {
	case FormatARGB32:
		if i, ok := img.(*extimage.BGRA); ok {
			if i.Rect.Dx() == width && i.Rect.Dy() == height && i.Stride == stride {
				s.SetData(i.Pix)
				return
			}
		}
		surfImg := s.GetImage().(*extimage.BGRA)
		draw.Draw(surfImg, surfImg.Bounds(), img, img.Bounds().Min, draw.Src)
		s.SetData(surfImg.Pix)

	case FormatRGB24:
		if i, ok := img.(*extimage.BGRN); ok {
			if i.Rect.Dx() == width && i.Rect.Dy() == height && i.Stride == stride {
				s.SetData(i.Pix)
				return
			}
		}
		surfImg := s.GetImage().(*extimage.BGRN)
		draw.Draw(surfImg, surfImg.Bounds(), img, img.Bounds().Min, draw.Src)
		s.SetData(surfImg.Pix)

	case FormatA8:
		if i, ok := img.(*image.Alpha); ok {
			if i.Rect.Dx() == width && i.Rect.Dy() == height && i.Stride == stride {
				s.SetData(i.Pix)
				return
			}
		}
		surfImg := s.GetImage().(*image.Alpha)
		draw.Draw(surfImg, surfImg.Bounds(), img, img.Bounds().Min, draw.Src)
		s.SetData(surfImg.Pix)

	case FormatA1:
		panic("Unsuppored surface format cairo.FormatA1")

	case FormatRGB16565:
		panic("Unsuppored surface format cairo.FORMAT_RGB16_565")

	case FormatRGB30:
		panic("Unsuppored surface format cairo.FORMAT_RGB30")

	case FormatInvalid:
		panic("Invalid surface format")

	default:
		panic("Unknown surface format")
	}
}
