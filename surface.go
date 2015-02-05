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

	"github.com/ungerik/go-cairo/extimage"
)

// Golang struct to hold both a cairo surface and a cairo context
type Surface struct {
	surface *C.cairo_surface_t
	context *C.cairo_t
}

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

func NewSurfaceFromPNG(filename string) *Surface {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	s := C.cairo_image_surface_create_from_png(cs)
	return &Surface{surface: s, context: C.cairo_create(s)}
}

func NewSurfaceFromImage(img image.Image) *Surface {
	var format Format
	switch img.(type) {
	case *image.Alpha, *image.Alpha16:
		format = FORMAT_A8
	case *extimage.BGRN, *image.Gray, *image.Gray16, *image.YCbCr:
		format = FORMAT_RGB24
	default:
		format = FORMAT_ARGB32
	}
	surface := NewSurface(format, img.Bounds().Dx(), img.Bounds().Dy())
	surface.SetImage(img)
	return surface
}

func NewPDFSurface(filename string, widthInPoints, heightInPoints float64, version PDFVersion) *Surface {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	s := C.cairo_pdf_surface_create(cs, C.double(widthInPoints), C.double(heightInPoints))
	C.cairo_pdf_surface_restrict_to_version(s, C.cairo_pdf_version_t(version))
	return &Surface{surface: s, context: C.cairo_create(s)}
}

func NewPSSurface(filename string, widthInPoints, heightInPoints float64, level PSLevel) *Surface {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	s := C.cairo_ps_surface_create(cs, C.double(widthInPoints), C.double(heightInPoints))
	C.cairo_ps_surface_restrict_to_level(s, C.cairo_ps_level_t(level))
	return &Surface{surface: s, context: C.cairo_create(s)}
}

func NewSVGSurface(filename string, widthInPoints, heightInPoints float64, version SVGVersion) *Surface {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	s := C.cairo_svg_surface_create(cs, C.double(widthInPoints), C.double(heightInPoints))
	C.cairo_svg_surface_restrict_to_version(s, C.cairo_svg_version_t(version))
	return &Surface{surface: s, context: C.cairo_create(s)}
}

func (self *Surface) Save() {
	C.cairo_save(self.context)
}

func (self *Surface) Restore() {
	C.cairo_restore(self.context)
}

func (self *Surface) PushGroup() {
	C.cairo_push_group(self.context)
}

func (self *Surface) PushGroupWithContent(content Content) {
	C.cairo_push_group_with_content(self.context, C.cairo_content_t(content))
}

func (self *Surface) PopGroup() (pattern *Pattern) {
	return &Pattern{C.cairo_pop_group(self.context)}
}

func (self *Surface) PopGroupToSource() {
	C.cairo_pop_group_to_source(self.context)
}

func (self *Surface) SetOperator(operator Operator) {
	C.cairo_set_operator(self.context, C.cairo_operator_t(operator))
}

func (self *Surface) SetSource(pattern *Pattern) {
	C.cairo_set_source(self.context, pattern.pattern)
}

func (self *Surface) SetSourceRGB(red, green, blue float64) {
	C.cairo_set_source_rgb(self.context, C.double(red), C.double(green), C.double(blue))
}

func (self *Surface) SetSourceRGBA(red, green, blue, alpha float64) {
	C.cairo_set_source_rgba(self.context, C.double(red), C.double(green), C.double(blue), C.double(alpha))
}

func (self *Surface) SetSourceSurface(surface *Surface, x, y float64) {
	C.cairo_set_source_surface(self.context, surface.surface, C.double(x), C.double(y))
}

func (self *Surface) SetTolerance(tolerance float64) {
	C.cairo_set_tolerance(self.context, C.double(tolerance))
}

func (self *Surface) SetAntialias(antialias Antialias) {
	C.cairo_set_antialias(self.context, C.cairo_antialias_t(antialias))
}

func (self *Surface) SetFillRule(fill_rule FillRule) {
	C.cairo_set_fill_rule(self.context, C.cairo_fill_rule_t(fill_rule))
}

func (self *Surface) SetLineWidth(width float64) {
	C.cairo_set_line_width(self.context, C.double(width))
}

func (self *Surface) SetLineCap(line_cap LineCap) {
	C.cairo_set_line_cap(self.context, C.cairo_line_cap_t(line_cap))
}

func (self *Surface) SetLineJoin(line_join LineJoin) {
	C.cairo_set_line_join(self.context, C.cairo_line_join_t(line_join))
}

func (self *Surface) SetDash(dashes []float64, num_dashes int, offset float64) {
	dashesp := (*C.double)(&dashes[0])
	C.cairo_set_dash(self.context, dashesp, C.int(num_dashes), C.double(offset))
}

func (self *Surface) SetMiterLimit(limit float64) {
	C.cairo_set_miter_limit(self.context, C.double(limit))
}

func (self *Surface) Translate(tx, ty float64) {
	C.cairo_translate(self.context, C.double(tx), C.double(ty))
}

func (self *Surface) Scale(sx, sy float64) {
	C.cairo_scale(self.context, C.double(sx), C.double(sy))
}

func (self *Surface) Rotate(angle float64) {
	C.cairo_rotate(self.context, C.double(angle))
}

func (self *Surface) Transform(matrix Matrix) {
	C.cairo_transform(self.context, matrix.cairo_matrix_t())
}

func (self *Surface) SetMatrix(matrix Matrix) {
	C.cairo_set_matrix(self.context, matrix.cairo_matrix_t())
}

func (self *Surface) IdentityMatrix() {
	C.cairo_identity_matrix(self.context)
}

func (self *Surface) UserToDevice(x, y float64) (float64, float64) {
	C.cairo_user_to_device(self.context, (*C.double)(&x), (*C.double)(&y))
	return x, y
}

func (self *Surface) UserToDeviceDistance(dx, dy float64) (float64, float64) {
	C.cairo_user_to_device_distance(self.context, (*C.double)(&dx), (*C.double)(&dy))
	return dx, dy
}

// path creation methods

func (self *Surface) NewPath() {
	C.cairo_new_path(self.context)
}

func (self *Surface) MoveTo(x, y float64) {
	C.cairo_move_to(self.context, C.double(x), C.double(y))
}

func (self *Surface) NewSubPath() {
	C.cairo_new_sub_path(self.context)
}

func (self *Surface) LineTo(x, y float64) {
	C.cairo_line_to(self.context, C.double(x), C.double(y))
}

func (self *Surface) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	C.cairo_curve_to(self.context,
		C.double(x1), C.double(y1),
		C.double(x2), C.double(y2),
		C.double(x3), C.double(y3))
}

func (self *Surface) Arc(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc(self.context,
		C.double(xc), C.double(yc),
		C.double(radius),
		C.double(angle1), C.double(angle2))
}

func (self *Surface) ArcNegative(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc_negative(self.context,
		C.double(xc), C.double(yc),
		C.double(radius),
		C.double(angle1), C.double(angle2))
}

func (self *Surface) RelMoveTo(dx, dy float64) {
	C.cairo_rel_move_to(self.context, C.double(dx), C.double(dy))
}

func (self *Surface) RelLineTo(dx, dy float64) {
	C.cairo_rel_line_to(self.context, C.double(dx), C.double(dy))
}

func (self *Surface) RelCurveTo(dx1, dy1, dx2, dy2, dx3, dy3 float64) {
	C.cairo_rel_curve_to(self.context,
		C.double(dx1), C.double(dy1),
		C.double(dx2), C.double(dy2),
		C.double(dx3), C.double(dy3))
}

func (self *Surface) Rectangle(x, y, width, height float64) {
	C.cairo_rectangle(self.context,
		C.double(x), C.double(y),
		C.double(width), C.double(height))
}

func (self *Surface) ClosePath() {
	C.cairo_close_path(self.context)
}

func (self *Surface) PathExtents() (left, top, right, bottom float64) {
	C.cairo_path_extents(self.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}

///////////////////////////////////////////////////////////////////////////////
// Painting methods

func (self *Surface) Paint() {
	C.cairo_paint(self.context)
}

func (self *Surface) PaintWithAlpha(alpha float64) {
	C.cairo_paint_with_alpha(self.context, C.double(alpha))
}

func (self *Surface) Mask(pattern Pattern) {
	C.cairo_mask(self.context, pattern.pattern)
}

func (self *Surface) MaskSurface(surface *Surface, surface_x, surface_y float64) {
	C.cairo_mask_surface(self.context, surface.surface, C.double(surface_x), C.double(surface_y))
}

func (self *Surface) Stroke() {
	C.cairo_stroke(self.context)
}

func (self *Surface) StrokePreserve() {
	C.cairo_stroke_preserve(self.context)
}

func (self *Surface) Fill() {
	C.cairo_fill(self.context)
}

func (self *Surface) FillPreserve() {
	C.cairo_fill_preserve(self.context)
}

func (self *Surface) CopyPage() {
	C.cairo_copy_page(self.context)
}

func (self *Surface) ShowPage() {
	C.cairo_show_page(self.context)
}

///////////////////////////////////////////////////////////////////////////////
// Insideness testing

func (self *Surface) InStroke(x, y float64) bool {
	return C.cairo_in_stroke(self.context, C.double(x), C.double(y)) != 0
}

func (self *Surface) InFill(x, y float64) bool {
	return C.cairo_in_fill(self.context, C.double(x), C.double(y)) != 0
}

///////////////////////////////////////////////////////////////////////////////
// Rectangular extents

func (self *Surface) StrokeExtents() (left, top, right, bottom float64) {
	C.cairo_stroke_extents(self.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}

func (self *Surface) FillExtents() (left, top, right, bottom float64) {
	C.cairo_fill_extents(self.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}

///////////////////////////////////////////////////////////////////////////////
// Clipping methods

func (self *Surface) ResetClip() {
	C.cairo_reset_clip(self.context)
}

func (self *Surface) Clip() {
	C.cairo_clip(self.context)
}

func (self *Surface) ClipPreserve() {
	C.cairo_clip_preserve(self.context)
}

func (self *Surface) ClipExtents() (left, top, right, bottom float64) {
	C.cairo_clip_extents(self.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}

func (self *Surface) ClipRectangleList() ([]Rectangle, Status) {
	list := C.cairo_copy_clip_rectangle_list(self.context)
	defer C.cairo_rectangle_list_destroy(list)
	rects := make([]Rectangle, int(list.num_rectangles))
	C.memcpy(unsafe.Pointer(&rects[0]), unsafe.Pointer(list.rectangles), C.size_t(list.num_rectangles*8))
	return rects, Status(list.status)
}

///////////////////////////////////////////////////////////////////////////////
// Font/Text methods

func (self *Surface) SelectFontFace(name string, font_slant_t, font_weight_t int) {
	s := C.CString(name)
	C.cairo_select_font_face(self.context, s, C.cairo_font_slant_t(font_slant_t), C.cairo_font_weight_t(font_weight_t))
	C.free(unsafe.Pointer(s))
}

func (self *Surface) SetFontSize(size float64) {
	C.cairo_set_font_size(self.context, C.double(size))
}

func (self *Surface) SetFontMatrix(matrix Matrix) {
	C.cairo_set_font_matrix(self.context, matrix.cairo_matrix_t())
}

func (self *Surface) SetFontOptions(fontOptions *FontOptions) {
	panic("not implemented") // todo
}

func (self *Surface) GetFontOptions() *FontOptions {
	panic("not implemented") // todo
	return nil
}

func (self *Surface) SetFontFace(fontFace *FontFace) {
	panic("not implemented") // todo
}

func (self *Surface) GetFontFace() *FontFace {
	panic("not implemented") // todo
	return nil
}

func (self *Surface) SetScaledFont(scaledFont *ScaledFont) {
	panic("not implemented") // todo
}

func (self *Surface) GetScaledFont() *ScaledFont {
	panic("not implemented") // todo
	return nil
}

func (self *Surface) ShowText(text string) {
	cs := C.CString(text)
	C.cairo_show_text(self.context, cs)
	C.free(unsafe.Pointer(cs))
}

func (self *Surface) ShowGlyphs(glyphs []Glyph) {
	panic("not implemented") // todo
}

func (self *Surface) ShowTextGlyphs(text string, glyphs []Glyph, clusters []TextCluster, flags TextClusterFlag) {
}

func (self *Surface) TextPath(text string) {
	cs := C.CString(text)
	C.cairo_text_path(self.context, cs)
	C.free(unsafe.Pointer(cs))
}

func (self *Surface) GlyphPath(glyphs []Glyph) {
	panic("not implemented") // todo
}

func (self *Surface) TextExtents(text string) *TextExtents {
	cte := C.cairo_text_extents_t{}
	cs := C.CString(text)
	C.cairo_text_extents(self.context, cs, &cte)
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

func (self *Surface) GlyphExtents(glyphs []Glyph) *TextExtents {
	panic("not implemented") // todo
	//C.cairo_text_extents
	return nil
}

func (self *Surface) FontExtents() *FontExtents {
	panic("not implemented") // todo
	//C.cairo_text_extents
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// Error status queries

func (self *Surface) Status() Status {
	return Status(C.cairo_status(self.context))
}

///////////////////////////////////////////////////////////////////////////////
// Backend device manipulation

///////////////////////////////////////////////////////////////////////////////
// Surface manipulation

func (self *Surface) CreateForRectangle(x, y, width, height float64) *Surface {
	return &Surface{
		context: self.context,
		surface: C.cairo_surface_create_for_rectangle(self.surface,
			C.double(x), C.double(y), C.double(width), C.double(height)),
	}
}

func (self *Surface) Finish() {
	C.cairo_surface_finish(self.surface)
}

func (self *Surface) Destroy() {
	C.cairo_surface_destroy(self.surface)
}

func (self *Surface) GetDevice() *Device {
	//C.cairo_surface_get_device
	panic("not implemented") // todo
	return nil
}

func (self *Surface) GetReferenceCount() int {
	return int(C.cairo_surface_get_reference_count(self.surface))
}

func (self *Surface) GetStatus() Status {
	return Status(C.cairo_surface_status(self.surface))
}

func (self *Surface) GetType() SurfaceType {
	return SurfaceType(C.cairo_surface_get_type(self.surface))
}

func (self *Surface) GetContent() Content {
	return Content(C.cairo_surface_get_content(self.surface))
}

func (self *Surface) WriteToPNG(filename string) {
	cs := C.CString(filename)
	C.cairo_surface_write_to_png(self.surface, cs)
	C.free(unsafe.Pointer(cs))
}

// Already implemented via context split context/surface?
// func (self *Surface) GetFontOptions() *FontOptions {
// 	// todo
// 	// C.cairo_surface_get_font_options (cairo_surface_t      *surface,				cairo_font_options_t *options);
// 	return nil
// }

func (self *Surface) Flush() {
	C.cairo_surface_flush(self.surface)
}

func (self *Surface) MarkDirty() {
	C.cairo_surface_mark_dirty(self.surface)
}

func (self *Surface) MarkDirtyRectangle(x, y, width, height int) {
	C.cairo_surface_mark_dirty_rectangle(self.surface,
		C.int(x), C.int(y), C.int(width), C.int(height))
}

func (self *Surface) SetDeviceOffset(x, y float64) {
	C.cairo_surface_set_device_offset(self.surface, C.double(x), C.double(y))
}

func (self *Surface) GetDeviceOffset() (x, y float64) {
	C.cairo_surface_get_device_offset(self.surface, (*C.double)(&x), (*C.double)(&y))
	return x, y
}

func (self *Surface) SetFallbackResolution(xPixelPerInch, yPixelPerInch float64) {
	C.cairo_surface_set_fallback_resolution(self.surface,
		C.double(xPixelPerInch), C.double(yPixelPerInch))
}

func (self *Surface) GetFallbackResolution() (xPixelPerInch, yPixelPerInch float64) {
	C.cairo_surface_get_fallback_resolution(self.surface,
		(*C.double)(&xPixelPerInch), (*C.double)(&yPixelPerInch))
	return xPixelPerInch, yPixelPerInch
}

// Already defined for context
// func (self *Surface) CopyPage() {
// 	C.cairo_surface_copy_page(self.surface)
// }

// func (self *Surface) ShowPage() {
// 	C.cairo_surface_show_page(self.surface)
// }

func (self *Surface) HasShowTextGlyphs() bool {
	return C.cairo_surface_has_show_text_glyphs(self.surface) != 0
}

// GetData returns a copy of the surfaces raw pixel data.
// This method also calls Flush.
func (self *Surface) GetData() []byte {
	self.Flush()
	dataPtr := C.cairo_image_surface_get_data(self.surface)
	if dataPtr == nil {
		panic("cairo.Surface.GetData(): can't access surface pixel data")
	}
	stride := C.cairo_image_surface_get_stride(self.surface)
	height := C.cairo_image_surface_get_height(self.surface)
	return C.GoBytes(unsafe.Pointer(dataPtr), stride*height)
}

// SetData sets the surfaces raw pixel data.
// This method also calls Flush and MarkDirty.
func (self *Surface) SetData(data []byte) {
	self.Flush()
	dataPtr := unsafe.Pointer(C.cairo_image_surface_get_data(self.surface))
	if dataPtr == nil {
		panic("cairo.Surface.SetData(): can't access surface pixel data")
	}
	stride := C.cairo_image_surface_get_stride(self.surface)
	height := C.cairo_image_surface_get_height(self.surface)
	if len(data) != int(stride*height) {
		panic("cairo.Surface.SetData(): invalid data size")
	}
	C.memcpy(dataPtr, unsafe.Pointer(&data[0]), C.size_t(stride*height))
	self.MarkDirty()
}

func (self *Surface) GetFormat() Format {
	return Format(C.cairo_image_surface_get_format(self.surface))
}

func (self *Surface) GetWidth() int {
	return int(C.cairo_image_surface_get_width(self.surface))
}

func (self *Surface) GetHeight() int {
	return int(C.cairo_image_surface_get_height(self.surface))
}

func (self *Surface) GetStride() int {
	return int(C.cairo_image_surface_get_stride(self.surface))
}

///////////////////////////////////////////////////////////////////////////////
// Pattern creation methods

///////////////////////////////////////////////////////////////////////////////
// image.Image methods

func (self *Surface) GetImage() image.Image {
	width := self.GetWidth()
	height := self.GetHeight()
	stride := self.GetStride()
	data := self.GetData()

	switch self.GetFormat() {
	case FORMAT_ARGB32:
		return &extimage.BGRA{
			Pix:    data,
			Stride: stride,
			Rect:   image.Rect(0, 0, width, height),
		}

	case FORMAT_RGB24:
		return &extimage.BGRN{
			Pix:    data,
			Stride: stride,
			Rect:   image.Rect(0, 0, width, height),
		}

	case FORMAT_A8:
		return &image.Alpha{
			Pix:    data,
			Stride: stride,
			Rect:   image.Rect(0, 0, width, height),
		}

	case FORMAT_A1:
		panic("Unsuppored surface format cairo.FORMAT_A1")

	case FORMAT_RGB16_565:
		panic("Unsuppored surface format cairo.FORMAT_RGB16_565")

	case FORMAT_RGB30:
		panic("Unsuppored surface format cairo.FORMAT_RGB30")

	case FORMAT_INVALID:
		panic("Invalid surface format")
	}
	panic("Unknown surface format")
}

// SetImage set the data from an image.Image with identical size.
func (self *Surface) SetImage(img image.Image) {
	width := self.GetWidth()
	height := self.GetHeight()
	stride := self.GetStride()

	switch self.GetFormat() {
	case FORMAT_ARGB32:
		if i, ok := img.(*extimage.BGRA); ok {
			if i.Rect.Dx() == width && i.Rect.Dy() == height && i.Stride == stride {
				self.SetData(i.Pix)
				return
			}
		}
		surfImg := self.GetImage().(*extimage.BGRA)
		draw.Draw(surfImg, surfImg.Bounds(), img, img.Bounds().Min, draw.Src)
		self.SetData(surfImg.Pix)

	case FORMAT_RGB24:
		if i, ok := img.(*extimage.BGRN); ok {
			if i.Rect.Dx() == width && i.Rect.Dy() == height && i.Stride == stride {
				self.SetData(i.Pix)
				return
			}
		}
		surfImg := self.GetImage().(*extimage.BGRN)
		draw.Draw(surfImg, surfImg.Bounds(), img, img.Bounds().Min, draw.Src)
		self.SetData(surfImg.Pix)

	case FORMAT_A8:
		if i, ok := img.(*image.Alpha); ok {
			if i.Rect.Dx() == width && i.Rect.Dy() == height && i.Stride == stride {
				self.SetData(i.Pix)
				return
			}
		}
		surfImg := self.GetImage().(*image.Alpha)
		draw.Draw(surfImg, surfImg.Bounds(), img, img.Bounds().Min, draw.Src)
		self.SetData(surfImg.Pix)

	case FORMAT_A1:
		panic("Unsuppored surface format cairo.FORMAT_A1")

	case FORMAT_RGB16_565:
		panic("Unsuppored surface format cairo.FORMAT_RGB16_565")

	case FORMAT_RGB30:
		panic("Unsuppored surface format cairo.FORMAT_RGB30")

	case FORMAT_INVALID:
		panic("Invalid surface format")

	default:
		panic("Unknown surface format")
	}
}
