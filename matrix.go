// +build !goci

package cairo

// #include <cairo/cairo.h>
import "C"

import (
	"unsafe"
)

type Matrix struct {
	Xx, Yx float64
	Xy, Yy float64
	X0, Y0 float64
}

func (self *Matrix) cairo_matrix_t() *C.cairo_matrix_t {
	return (*C.cairo_matrix_t)(unsafe.Pointer(self))
}

func (self *Matrix) InitIdendity() {
	C.cairo_matrix_init_identity(self.cairo_matrix_t())
}

func (self *Matrix) InitTranslate(tx, ty float64) {
	C.cairo_matrix_init_translate(self.cairo_matrix_t(), C.double(tx), C.double(ty))
}

func (self *Matrix) InitScale(sx, sy float64) {
	C.cairo_matrix_init_scale(self.cairo_matrix_t(), C.double(sx), C.double(sy))
}

func (self *Matrix) InitRotate(radians float64) {
	C.cairo_matrix_init_rotate(self.cairo_matrix_t(), C.double(radians))
}

func (self *Matrix) Translate(tx, ty float64) {
	C.cairo_matrix_translate(self.cairo_matrix_t(), C.double(tx), C.double(ty))
}

func (self *Matrix) Scale(sx, sy float64) {
	C.cairo_matrix_scale(self.cairo_matrix_t(), C.double(sx), C.double(sy))
}

func (self *Matrix) Rotate(radians float64) {
	C.cairo_matrix_rotate(self.cairo_matrix_t(), C.double(radians))
}

func (self *Matrix) Invert() {
	C.cairo_matrix_invert(self.cairo_matrix_t())
}

func (self *Matrix) Multiply(a, b Matrix) {
	C.cairo_matrix_multiply(self.cairo_matrix_t(), a.cairo_matrix_t(), b.cairo_matrix_t())
}

func (self *Matrix) TransformDistance(dx, dy float64) (float64, float64) {
	C.cairo_matrix_transform_distance(self.cairo_matrix_t(),
		(*C.double)(unsafe.Pointer(&dx)), (*C.double)(unsafe.Pointer(&dy)))
	return dx, dy
}

func (self *Matrix) TransformPoint(x, y float64) (float64, float64) {
	C.cairo_matrix_transform_point(self.cairo_matrix_t(),
		(*C.double)(unsafe.Pointer(&x)), (*C.double)(unsafe.Pointer(&y)))
	return x, y
}
