// package cairo contains constants for use in cairo
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// TextClusterFlag cairo_text_cluster_flag_t
type TextClusterFlag int

// TextClusterFlag constants
const (
	TextClusterFlagBackward TextClusterFlag = 0x00000001
)

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

// cairo_subpixel_order_t values
const (
	SubpixelOrderDefault = iota
	SubpixelOrderRGB
	SubpixelOrderBGR
	SubpixelOrderVRGB
	SubpixelOrderVBGR
)

// cairo_hint_style_t values
const (
	HintStyleDefault = iota
	HintStyleNone
	HintStyleSlight
	HintStyleMedium
	HintStyleFull
)

// cairo_hint_metrics_t values
const (
	HintMetricsDefault = iota
	HintMetricsOff
	HintMetricsOn
)

// FontType cairo_font_type_t
type FontType int

// FontType constants
const (
	FontTypeToy FontType = iota
	FontTypeFt
	FontTypeWin32
	FontTypeQuartz
	FontTypeUser
)

// TextCluster cairo_text_cluster_t
type TextCluster struct {
	// todo
}

// TextExtents cairo_text_extents_t
type TextExtents struct {
	Xbearing float64
	Ybearing float64
	Width    float64
	Height   float64
	Xadvance float64
	Yadvance float64
}

// FontExtents cairo_font_extents_t
type FontExtents struct {
	// todo
}

// FontFace cairo_font_face_t
type FontFace struct {
	// todo
}

// FontOptions cairo_font_options_t
type FontOptions struct {
	// todo
}

// ScaledFont cairo_scaled_font_t
type ScaledFont struct {
	// todo
}

// Glyph cairo_glyph_t
type Glyph struct {
	// todo
}
