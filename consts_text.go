// Package cairo wraps the c cairographics library.
package cairo

// cairo_font_slant_t constants
const (
	FontSlantNormal = iota
	FontSlantItalic
	FontSlantOblique
)

// cairo_font_weight_t values
const (
	FontWeightNormal = iota
	FontWeightBold
)

// TextExtents cairo_text_extents_t
type TextExtents struct {
	Xbearing float64
	Ybearing float64
	Width    float64
	Height   float64
	Xadvance float64
	Yadvance float64
}
