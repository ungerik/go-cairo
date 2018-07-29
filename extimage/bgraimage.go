package extimage

import (
	"image"
	"image/color"
)

// NewBGRA returns a new BGRA with the given bounds.
func NewBGRA(r image.Rectangle) *BGRA {
	w, h := r.Dx(), r.Dy()
	buf := make([]uint8, 4*w*h)
	return &BGRA{Pix: buf, Stride: 4 * w, Rect: r}
}

// BGRA is an in-memory image whose At method returns BGRAColor values.
type BGRA struct {
	// Pix holds the image's pixels, in B, G, R, A order on small endian systems
	// and A, R, G, B on big endian systems.
	// See http://cairographics.org/manual/cairo-Image-Surfaces.html#cairo-format-t
	// The pixel at (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

// ColorModel ...
func (img *BGRA) ColorModel() color.Model {
	return BGRAColorModel
}

// Bounds ...
func (img *BGRA) Bounds() image.Rectangle {
	return img.Rect
}

// At ...
func (img *BGRA) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(img.Rect)) {
		return BGRAColor{}
	}
	i := img.PixOffset(x, y)
	if littleEndian {
		return BGRAColor{B: img.Pix[i+0], G: img.Pix[i+1], R: img.Pix[i+2], A: img.Pix[i+3]}
	}
	return BGRAColor{A: img.Pix[i+0], R: img.Pix[i+1], G: img.Pix[i+2], B: img.Pix[i+3]}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (img *BGRA) PixOffset(x, y int) int {
	return (y-img.Rect.Min.Y)*img.Stride + (x-img.Rect.Min.X)*4
}

// Set ...
func (img *BGRA) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(img.Rect)) {
		return
	}
	i := img.PixOffset(x, y)
	c1 := BGRAColorModel.Convert(c).(BGRAColor)
	if littleEndian {
		img.Pix[i+0] = c1.B
		img.Pix[i+1] = c1.G
		img.Pix[i+2] = c1.R
		img.Pix[i+3] = c1.A
	} else {
		img.Pix[i+0] = c1.A
		img.Pix[i+1] = c1.R
		img.Pix[i+2] = c1.G
		img.Pix[i+3] = c1.B
	}
}

// BGRAColorModel ...
var BGRAColorModel = color.ModelFunc(
	func(c color.Color) color.Color {
		if _, ok := c.(BGRAColor); ok {
			return c
		}
		r, g, b, a := c.RGBA()
		return BGRAColor{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: uint8(a >> 8)}
	},
)

// BGRAColor represents a traditional 32-bit alpha-premultiplied color,
// having 8 bits for each of alpha, red, green and blue.
type BGRAColor struct {
	B, G, R, A uint8
}

// RGBA ...
func (c BGRAColor) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8

	g = uint32(c.G)
	g |= g << 8

	b = uint32(c.B)
	b |= b << 8

	a = uint32(c.A)
	a |= a << 8

	return r, g, b, a
}
