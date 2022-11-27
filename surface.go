// Package cairo wraps the c cairographics library.
package cairo

// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"unsafe"
)

// Surface represents a cairo surface
type Surface struct {
	surface *C.cairo_surface_t
}

// NewSurface creates a new cairo surface.
func NewSurface(format Format, width, height int) *Surface {
	return &Surface{
		C.cairo_image_surface_create(C.cairo_format_t(format), C.int(width), C.int(height)),
	}
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

	surface := &Surface{
		surface: surfaceNative,
	}

	return surface, StatusSuccess
}

// Native returns a pointer to the native cairo surface
func (s *Surface) Native() uintptr {
	return uintptr(unsafe.Pointer(s.surface))
}

// CreateForRectangle tbd
func (s *Surface) CreateForRectangle(x, y, width, height float64) *Surface {
	return &Surface{
		surface: C.cairo_surface_create_for_rectangle(s.surface,
			C.double(x), C.double(y), C.double(width), C.double(height)),
	}
}

// Finish tbd
func (s *Surface) Finish() {
	C.cairo_surface_finish(s.surface)
}

// Destroy tbd
func (s *Surface) Destroy() {
	C.cairo_surface_destroy(s.surface)
}

// GetDevice tbd
func (s *Surface) GetDevice() *Device {
	//C.cairo_surface_get_device
	panic("not implemented") // todo
}

// GetReferenceCount tbd
func (s *Surface) GetReferenceCount() int {
	return int(C.cairo_surface_get_reference_count(s.surface))
}

// GetStatus tbd
func (s *Surface) GetStatus() Status {
	return Status(C.cairo_surface_status(s.surface))
}

// GetType tbd
func (s *Surface) GetType() SurfaceType {
	return SurfaceType(C.cairo_surface_get_type(s.surface))
}

// GetContent tbd
func (s *Surface) GetContent() Content {
	return Content(C.cairo_surface_get_content(s.surface))
}

// WriteToPNG tbd
func (s *Surface) WriteToPNG(filename string) Status {

	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))

	return Status(C.cairo_surface_write_to_png(s.surface, cs))
}

// Flush tbd
func (s *Surface) Flush() {
	C.cairo_surface_flush(s.surface)
}

// MarkDirty tbd
func (s *Surface) MarkDirty() {
	C.cairo_surface_mark_dirty(s.surface)
}

// MarkDirtyRectangle tbd
func (s *Surface) MarkDirtyRectangle(x, y, width, height int) {
	C.cairo_surface_mark_dirty_rectangle(s.surface,
		C.int(x), C.int(y), C.int(width), C.int(height))
}

// SetDeviceOffset tbd
func (s *Surface) SetDeviceOffset(x, y float64) {
	C.cairo_surface_set_device_offset(s.surface, C.double(x), C.double(y))
}

// GetDeviceOffset tbd
func (s *Surface) GetDeviceOffset() (x, y float64) {
	C.cairo_surface_get_device_offset(s.surface, (*C.double)(&x), (*C.double)(&y))
	return x, y
}

// SetFallbackResolution tbd
func (s *Surface) SetFallbackResolution(xPixelPerInch, yPixelPerInch float64) {
	C.cairo_surface_set_fallback_resolution(s.surface,
		C.double(xPixelPerInch), C.double(yPixelPerInch))
}

// GetFallbackResolution tbd
func (s *Surface) GetFallbackResolution() (xPixelPerInch, yPixelPerInch float64) {
	C.cairo_surface_get_fallback_resolution(s.surface,
		(*C.double)(&xPixelPerInch), (*C.double)(&yPixelPerInch))
	return xPixelPerInch, yPixelPerInch
}

// HasShowTextGlyphs tbd
func (s *Surface) HasShowTextGlyphs() bool {
	return C.cairo_surface_has_show_text_glyphs(s.surface) != 0
}

// GetData returns a copy of the surfaces raw pixel data.
// This method also calls Flush.
// GetData tbd
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
// SetData tbd
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

// GetFormat tbd
func (s *Surface) GetFormat() Format {
	return Format(C.cairo_image_surface_get_format(s.surface))
}

// GetWidth tbd
func (s *Surface) GetWidth() int {
	return int(C.cairo_image_surface_get_width(s.surface))
}

// GetHeight tbd
func (s *Surface) GetHeight() int {
	return int(C.cairo_image_surface_get_height(s.surface))
}

// GetStride tbd
func (s *Surface) GetStride() int {
	return int(C.cairo_image_surface_get_stride(s.surface))
}
