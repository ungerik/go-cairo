package extimage

import (
	"image"
	"image/color"
)

// NewRGBA returns a new RGBA with the given bounds.
func NewARGB(r image.Rectangle) *ARGB {
	w, h := r.Dx(), r.Dy()
	buf := make([]uint8, 4*w*h)
	return &ARGB{buf, 4 * w, r}
}

// ARGB is an in-memory image whose At method returns ARGBColor values.
type ARGB struct {
	// Pix holds the image's pixels, in A, R, G, B order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (self *ARGB) ColorModel() color.Model {
	return ARGBColorModel
}

func (self *ARGB) Bounds() image.Rectangle {
	return self.Rect
}

func (self *ARGB) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(self.Rect)) {
		return ARGBColor{}
	}
	i := self.PixOffset(x, y)
	return ARGBColor{self.Pix[i+0], self.Pix[i+1], self.Pix[i+2], self.Pix[i+3]}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (self *ARGB) PixOffset(x, y int) int {
	return (y-self.Rect.Min.Y)*self.Stride + (x-self.Rect.Min.X)*4
}

func (self *ARGB) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(self.Rect)) {
		return
	}
	i := self.PixOffset(x, y)
	c1 := ARGBColorModel.Convert(c).(ARGBColor)
	self.Pix[i+0] = c1.A
	self.Pix[i+1] = c1.R
	self.Pix[i+2] = c1.G
	self.Pix[i+3] = c1.B
}

var ARGBColorModel = color.ModelFunc(
	func(c color.Color) color.Color {
		if _, ok := c.(ARGBColor); ok {
			return c
		}
		r, g, b, a := c.RGBA()
		return ARGBColor{uint8(a >> 8), uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
	},
)

// RGBA represents a traditional 32-bit alpha-premultiplied color,
// having 8 bits for each of alpha, red, green and blue.
type ARGBColor struct {
	A, R, G, B uint8
}

func (c ARGBColor) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}
