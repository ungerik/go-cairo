package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// Status cairo_status_t
type Status int

func (s Status) String() string {
	return C.GoString(C.cairo_status_to_string(C.cairo_status_t(s)))
}

// Status constants
const (
	StatusSuccess Status = iota
	StatusNoMemory
	StatusInvalidRestore
	StatusInvalidPopGroup
	StatusNoCurrentPoint
	StatusInvalidMatrix
	StatusInvalidStatus
	StatusNullPointer
	StatusInvalidString
	StatusInvalidPathData
	StatusReadError
	StatusWriteError
	StatusSurfaceFinished
	StatusSurfaceTypeMismatch
	StatusPatternTypeMismatch
	StatusInvalidContent
	StatusInvalidFormat
	StatusInvalidVisual
	StatusFileNotFound
	StatusInvalidDash
	StatusInvalidDscComment
	StatusInvalidIndex
	StatusClipNotRepresentable
	StatusTempFileError
	StatusInvalidStride
	StatusFontTypeMismatch
	StatusUserFontImmutable
	StatusUserFontError
	StatusNegativeCount
	StatusInvalidClusters
	StatusInvalidSlant
	StatusInvalidWeight
	StatusInvalidSize
	StatusUserFontNotImplemented
	statusDeviceTypeMismatch
	StatusDeviceError
)

// Content cairo_content_t
type Content int

// Content constants
const (
	ContentColor      Content = 0x1000
	ContentAlpha      Content = 0x2000
	ContentColorAlpha Content = 0x3000
)

// Operator cairo_operator_t
type Operator int

// Operator constants
const (
	OperatorClear = iota

	OperatorSource
	OperatorOver
	OperatorIn
	OperatorOut
	OperatorAtop

	OperatorDest
	OperatorDestOver
	OperatorDestIn
	OperatorDestOut
	OperatorDestAtop

	OperatorXor
	OperatorAdd
	OperatorSaturate

	OperatorMultiply
	OperatorScreen
	OperatorOverlay
	OperatorDarken
	OperatorLighten
	OperatorColorDodge
	OperatorColorBurn
	OperatorHardLight
	OperatorSoftLight
	OperatorDifference
	OperatorExclusion
	OperatorHslHue
	OperatorHslSaturation
	OperatorHslColor
	OperatorHslLuminosity
)

// Antialias cairo_antialias_t
type Antialias int

// Antialias constants
const (
	AntialiasDefault Antialias = iota
	AntialiasNone
	AntialiasGray
	AntialiasSubpixel
)

// FillRule cairo_fill_rule_t
type FillRule int

// FillRule constants
const (
	FillRuleWinding FillRule = iota
	FillRuleEvenOdd
)

// LineCap cairo_line_cap_t
type LineCap int

// LineCap constants
const (
	LineCapButt LineCap = iota
	LineCapRound
	LineCapSquare
)

// LineJoin cairo_line_cap_join_t
type LineJoin int

// LineJoin constants
const (
	LineJoinMiter LineJoin = iota
	LineJoinRound
	LineJoinBevel
)

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

// PathDataType cairo_path_data_type_t
type PathDataType int

// PathDataType constants
const (
	PathMoveTo PathDataType = iota
	PathLineTo
	PathCurveTo
	PathClosePath
)

// SurfaceType cairo_surface_type_t
type SurfaceType int

// SurfaceType constants
const (
	SurfaceTypeImage SurfaceType = iota
	SurfaceTypePDF
	SurfaceTypePS
	SurfaceTypeXlib
	SurfaceTypeXcb
	SurfaceTypeGlitz
	SurfaceTypeQuartz
	SurfaceTypeWin32
	SurfaceTypeBeos
	SurfaceTypeDirectfb
	SurfaceTypeSVG
	SurfaceTypeOS2
	SurfaceTypeWin32PrintinG
	SurfaceTypeQuartzImage
	SurfaceTypeScript
	SurfaceTypeQt
	SurfaceTypeRecording
	SurfaceTypeVg
	SurfaceTypeGl
	SurfaceTypeDrm
	SurfaceTypeTee
	SurfaceTypeXML
	SurfaceTypeSkia
	SurfaceTypeSubsurface
)

// Format cairo_format_t
type Format int

// StrideForWidth provides a stride value that will respect all alignment requirements of the accelerated image-rendering code within cairo
func (f Format) StrideForWidth(width int) int {
	return int(C.cairo_format_stride_for_width(C.cairo_format_t(f), C.int(width)))
}

// Format constants
const (
	FormatInvalid  Format = -1
	FormatARGB32   Format = 0
	FormatRGB24    Format = 1
	FormatA8       Format = 2
	FormatA1       Format = 3
	FormatRGB16565 Format = 4
	FormatRGB30    Format = 5
)

// Extent cairo_extent_t
type Extent int

// Extent constants
const (
	ExtendNone Extent = iota
	ExtendRepeat
	ExtendReflect
	ExtendPad
)

// Filter cairo_filter_t
type Filter int

// Filter constants
const (
	CairoFilterFast Filter = iota
	CairoFilterGood
	CairoFilterBest
	CairoFilterNearest
	CairoFilterBilinear
	CairoFilterGaussian
)

// RegionOverlap Used as the return value for cairo_region_contains_rectangle().
type RegionOverlap int

// RegionOverlap constants
const (
	RegionOverlapIn RegionOverlap = iota
	RegionOverlapOut
	RegionOverlapPart
)

// DeviceType cairo_device_t
type DeviceType int

// DeviceType constants
const (
	DeviceTypeDrm DeviceType = iota
	DeviceTypeGl
	DeviceTypeScript
	DeviceTypeXcb
	DeviceTypeXlib
	DeviceTypeXML
)

// MimeType constants
const (
	MimeTypeJPEG = "image/jpeg"
	MimeTypePNG  = "image/png"
	MimeTypeJP2  = "image/jp2"
	MimeTypeURI  = "text/x-uri"
)

// PDFVersion version of PDF created
type PDFVersion int

func (v PDFVersion) String() string {
	return C.GoString(C.cairo_pdf_version_to_string(C.cairo_pdf_version_t(v)))
}

// PDFVersion constants
const (
	PDFVersion14 PDFVersion = iota
	PDFVersion15
)

// PSLevel PostScript level
type PSLevel int

func (p PSLevel) String() string {
	return C.GoString(C.cairo_ps_level_to_string(C.cairo_ps_level_t(p)))
}

// PSLevel constants
const (
	PSLevel2 PSLevel = iota
	PSLevel3
)

// SVGVersion SVG Version
type SVGVersion int

func (v SVGVersion) String() string {
	return C.GoString(C.cairo_svg_version_to_string(C.cairo_svg_version_t(v)))
}

// SVGVersion constants
const (
	SVGVersion11 SVGVersion = iota
	SVGVersion12
)

// Rectangle rectangle struct
type Rectangle struct {
	X, Y          float64
	Width, Height float64
}

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

// Device cairo_device_t
type Device struct {
}

// Version cairo_version
func Version() int {
	return int(C.cairo_version())
}

// VersionString cairo_version_string
func VersionString() string {
	return C.GoString(C.cairo_version_string())
}
