package extimage

import (
	"image"
	"image/color"
)

// NewRGBA returns a new RGB with the given bounds.
func NewRGB(r image.Rectangle) *RGB {
	w, h := r.Dx(), r.Dy()
	stride := 3 * w
	if stride&3 != 0 {
		stride += 4 - stride&3
	}
	buf := make([]uint8, h*stride)
	return &RGB{buf, stride, r}
}

// RGB is an in-memory image whose At method returns RGBColor values.
type RGB struct {
	// Pix holds the image's pixels, in R, G, B order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*3].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (self *RGB) ColorModel() color.Model {
	return RGBColorModel
}

func (self *RGB) Bounds() image.Rectangle {
	return self.Rect
}

func (self *RGB) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(self.Rect)) {
		return RGBColor{}
	}
	i := self.PixOffset(x, y)
	return RGBColor{self.Pix[i+0], self.Pix[i+1], self.Pix[i+2]}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (self *RGB) PixOffset(x, y int) int {
	return (y-self.Rect.Min.Y)*self.Stride + (x-self.Rect.Min.X)*3
}

func (self *RGB) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(self.Rect)) {
		return
	}
	i := self.PixOffset(x, y)
	c1 := RGBColorModel.Convert(c).(RGBColor)
	self.Pix[i+0] = c1.R
	self.Pix[i+1] = c1.G
	self.Pix[i+2] = c1.B
}

var RGBColorModel = color.ModelFunc(
	func(c color.Color) color.Color {
		if _, ok := c.(RGBColor); ok {
			return c
		}
		r, g, b, _ := c.RGBA()
		return RGBColor{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
	},
)

// RGBA represents a traditional 32-bit alpha-premultiplied color,
// having 8 bits for each of alpha, red, green and blue.
type RGBColor struct {
	R, G, B uint8
}

func (c RGBColor) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = 255
	return
}
