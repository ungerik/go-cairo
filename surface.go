// Package cairo wraps the c cairographics library.
package cairo

// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"errors"
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
func NewSurfaceFromPNG(filename string) (*Surface, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))

	surfaceNative := C.cairo_image_surface_create_from_png(cstr)
	status := Status(C.cairo_surface_status(surfaceNative))
	if status != StatusSuccess {
		return nil, errors.New(status.String())
	}

	surface := &Surface{
		surface: surfaceNative,
	}

	return surface, nil
}

// Finish finishes the surface. Further drawing operations will fail.
func (s *Surface) Finish() {
	C.cairo_surface_finish(s.surface)
}

// Destroy destroys the surface when no more references exist to it, freeing memory.
func (s *Surface) Destroy() {
	C.cairo_surface_destroy(s.surface)
}

// GetReferenceCount gets the number of objects that are holding a reference to this surface.
func (s *Surface) GetReferenceCount() int {
	return int(C.cairo_surface_get_reference_count(s.surface))
}

// GetStatus gets the status of the surface based on the last operation.
func (s *Surface) GetStatus() Status {
	return Status(C.cairo_surface_status(s.surface))
}

// GetContent gets the content type of the surface.
func (s *Surface) GetContent() Content {
	return Content(C.cairo_surface_get_content(s.surface))
}

// WriteToPNG saves the surface to an external png file
func (s *Surface) WriteToPNG(filename string) error {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	status := Status(C.cairo_surface_write_to_png(s.surface, cs))
	if status != StatusSuccess {
		return errors.New(status.String())
	}
	return nil
}

// Flush finished any pending drawing operations on the surface.
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
func (s *Surface) GetData() ([]byte, error) {
	s.Flush()
	dataPtr := C.cairo_image_surface_get_data(s.surface)
	if dataPtr == nil {
		return nil, errors.New("cairo.Surface.GetData(): can't access surface pixel data")
	}
	stride := C.cairo_image_surface_get_stride(s.surface)
	height := C.cairo_image_surface_get_height(s.surface)
	return C.GoBytes(unsafe.Pointer(dataPtr), stride*height), nil
}

// SetData sets the surfaces raw pixel data.
// This method also calls Flush and MarkDirty.
func (s *Surface) SetData(data []byte) error {
	s.Flush()
	dataPtr := unsafe.Pointer(C.cairo_image_surface_get_data(s.surface))
	if dataPtr == nil {
		return errors.New("cairo.Surface.SetData(): can't access surface pixel data")
	}
	stride := C.cairo_image_surface_get_stride(s.surface)
	height := C.cairo_image_surface_get_height(s.surface)
	if len(data) != int(stride*height) {
		return errors.New("cairo.Surface.SetData(): invalid data size")
	}
	C.memcpy(dataPtr, unsafe.Pointer(&data[0]), C.size_t(stride*height))
	s.MarkDirty()
	return nil
}

// GetFormat returns the format of the surface.
func (s *Surface) GetFormat() Format {
	return Format(C.cairo_image_surface_get_format(s.surface))
}

// GetWidth returns the width of the surface.
func (s *Surface) GetWidth() int {
	return int(C.cairo_image_surface_get_width(s.surface))
}

// GetHeight returns the height of the surface.
func (s *Surface) GetHeight() int {
	return int(C.cairo_image_surface_get_height(s.surface))
}

// GetStride returns the stride of the surface.
func (s *Surface) GetStride() int {
	return int(C.cairo_image_surface_get_stride(s.surface))
}
