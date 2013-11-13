package extimage

import (
	"image"
	"image/color"
)

// NewBGRN returns a new BGRN with the given bounds.
func NewBGRN(r image.Rectangle) *BGRN {
	w, h := r.Dx(), r.Dy()
	buf := make([]uint8, 4*w*h)
	return &BGRN{Pix: buf, Stride: 4 * w, Rect: r}
}

// BGRN is an in-memory image whose At method returns BGRNColor values.
type BGRN struct {
	// Pix holds the image's pixels, in B, G, R order on small endian systems
	// and R, G, B on big endian systems.
	// See http://cairographics.org/manual/cairo-Image-Surfaces.html#cairo-format-t
	// The pixel at (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (self *BGRN) ColorModel() color.Model {
	return BGRNColorModel
}

func (self *BGRN) Bounds() image.Rectangle {
	return self.Rect
}

func (self *BGRN) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(self.Rect)) {
		return BGRNColor{}
	}
	i := self.PixOffset(x, y)
	if littleEndian {
		return BGRNColor{B: self.Pix[i+0], G: self.Pix[i+1], R: self.Pix[i+2]}
	} else {
		return BGRNColor{R: self.Pix[i+1], G: self.Pix[i+2], B: self.Pix[i+3]}
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (self *BGRN) PixOffset(x, y int) int {
	return (y-self.Rect.Min.Y)*self.Stride + (x-self.Rect.Min.X)*3
}

func (self *BGRN) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(self.Rect)) {
		return
	}
	i := self.PixOffset(x, y)
	c1 := BGRNColorModel.Convert(c).(BGRNColor)
	if littleEndian {
		self.Pix[i+0] = c1.B
		self.Pix[i+1] = c1.G
		self.Pix[i+2] = c1.R
	} else {
		self.Pix[i+1] = c1.R
		self.Pix[i+2] = c1.G
		self.Pix[i+3] = c1.B
	}
}

var BGRNColorModel = color.ModelFunc(
	func(c color.Color) color.Color {
		if _, ok := c.(BGRNColor); ok {
			return c
		}
		r, g, b, _ := c.RGBA()
		return BGRNColor{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8)}
	},
)

type BGRNColor struct {
	B, G, R uint8
}

func (c BGRNColor) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8

	g = uint32(c.G)
	g |= g << 8

	b = uint32(c.B)
	b |= b << 8

	return r, g, b, 0xffff
}
