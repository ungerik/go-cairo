// +build !goci

package cairo

// #cgo pkg-config: cairo, cairo-ft
// #include <cairo/cairo-ft.h>
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// cairo_status_t
type Status int

func (self Status) String() string {
	return C.GoString(C.cairo_status_to_string(C.cairo_status_t(self)))
}

const (
	STATUS_SUCCESS Status = iota
	STATUS_NO_MEMORY
	STATUS_INVALID_RESTORE
	STATUS_INVALID_POP_GROUP
	STATUS_NO_CURRENT_POINT
	STATUS_INVALID_MATRIX
	STATUS_INVALID_STATUS
	STATUS_NULL_POINTER
	STATUS_INVALID_STRING
	STATUS_INVALID_PATH_DATA
	STATUS_READ_ERROR
	STATUS_WRITE_ERROR
	STATUS_SURFACE_FINISHED
	STATUS_SURFACE_TYPE_MISMATCH
	STATUS_PATTERN_TYPE_MISMATCH
	STATUS_INVALID_CONTENT
	STATUS_INVALID_FORMAT
	STATUS_INVALID_VISUAL
	STATUS_FILE_NOT_FOUND
	STATUS_INVALID_DASH
	STATUS_INVALID_DSC_COMMENT
	STATUS_INVALID_INDEX
	STATUS_CLIP_NOT_REPRESENTABLE
	STATUS_TEMP_FILE_ERROR
	STATUS_INVALID_STRIDE
	STATUS_FONT_TYPE_MISMATCH
	STATUS_USER_FONT_IMMUTABLE
	STATUS_USER_FONT_ERROR
	STATUS_NEGATIVE_COUNT
	STATUS_INVALID_CLUSTERS
	STATUS_INVALID_SLANT
	STATUS_INVALID_WEIGHT
	STATUS_INVALID_SIZE
	STATUS_USER_FONT_NOT_IMPLEMENTED
	STATUS_DEVICE_TYPE_MISMATCH
	STATUS_DEVICE_ERROR
)

// cairo_content_t
type Content int

const (
	CONTENT_COLOR       Content = 0x1000
	CONTENT_ALPHA       Content = 0x2000
	CONTENT_COLOR_ALPHA Content = 0x3000
)

// cairo_operator_t
type Operator int

const (
	OPERATOR_CLEAR = iota

	OPERATOR_SOURCE
	OPERATOR_OVER
	OPERATOR_IN
	OPERATOR_OUT
	OPERATOR_ATOP

	OPERATOR_DEST
	OPERATOR_DEST_OVER
	OPERATOR_DEST_IN
	OPERATOR_DEST_OUT
	OPERATOR_DEST_ATOP

	OPERATOR_XOR
	OPERATOR_ADD
	OPERATOR_SATURATE

	OPERATOR_MULTIPLY
	OPERATOR_SCREEN
	OPERATOR_OVERLAY
	OPERATOR_DARKEN
	OPERATOR_LIGHTEN
	OPERATOR_COLOR_DODGE
	OPERATOR_COLOR_BURN
	OPERATOR_HARD_LIGHT
	OPERATOR_SOFT_LIGHT
	OPERATOR_DIFFERENCE
	OPERATOR_EXCLUSION
	OPERATOR_HSL_HUE
	OPERATOR_HSL_SATURATION
	OPERATOR_HSL_COLOR
	OPERATOR_HSL_LUMINOSITY
)

// cairo_antialias_t
type Antialias int

const (
	ANTIALIAS_DEFAULT Antialias = iota
	ANTIALIAS_NONE
	ANTIALIAS_GRAY
	ANTIALIAS_SUBPIXEL
)

// cairo_svg_unit_t
type SVGUnit int

const (
	CAIRO_SVG_UNIT_USER = iota
	CAIRO_SVG_UNIT_EM
	CAIRO_SVG_UNIT_EX
	CAIRO_SVG_UNIT_PX
	CAIRO_SVG_UNIT_IN
	CAIRO_SVG_UNIT_CM
	CAIRO_SVG_UNIT_MM
	CAIRO_SVG_UNIT_PT
	CAIRO_SVG_UNIT_PC
	CAIRO_SVG_UNIT_PERCENT
)

// cairo_fill_rule_t
type FillRule int

const (
	FILL_RULE_WINDING FillRule = iota
	FILL_RULE_EVEN_ODD
)

// cairo_line_cap_t
type LineCap int

const (
	LINE_CAP_BUTT LineCap = iota
	LINE_CAP_ROUND
	LINE_CAP_SQUARE
)

// cairo_line_cap_join_t
type LineJoin int

const (
	LINE_JOIN_MITER LineJoin = iota
	LINE_JOIN_ROUND
	LINE_JOIN_BEVEL
)

// cairo_text_cluster_flag_t
type TextClusterFlag int

const (
	TEXT_CLUSTER_FLAG_BACKWARD TextClusterFlag = 0x00000001
)

// cairo_font_slant_t values
const (
	FONT_SLANT_NORMAL = iota
	FONT_SLANT_ITALIC
	FONT_SLANT_OBLIQUE
)

// cairo_font_weight_t values
const (
	FONT_WEIGHT_NORMAL = iota
	FONT_WEIGHT_BOLD
)

// cairo_subpixel_order_t values
const (
	SUBPIXEL_ORDER_DEFAULT = iota
	SUBPIXEL_ORDER_RGB
	SUBPIXEL_ORDER_BGR
	SUBPIXEL_ORDER_VRGB
	SUBPIXEL_ORDER_VBGR
)

// cairo_hint_style_t values
const (
	HINT_STYLE_DEFAULT = iota
	HINT_STYLE_NONE
	HINT_STYLE_SLIGHT
	HINT_STYLE_MEDIUM
	HINT_STYLE_FULL
)

// cairo_hint_metrics_t values
const (
	HINT_METRICS_DEFAULT = iota
	HINT_METRICS_OFF
	HINT_METRICS_ON
)

// cairo_font_type_t
type FontType int

const (
	FONT_TYPE_TOY FontType = iota
	FONT_TYPE_FT
	FONT_TYPE_WIN32
	FONT_TYPE_QUARTZ
	FONT_TYPE_USER
)

// cairo_path_data_type_t
type PathDataType int

const (
	PATH_MOVE_TO PathDataType = iota
	PATH_LINE_TO
	PATH_CURVE_TO
	PATH_CLOSE_PATH
)

// cairo_surface_type_t
type SurfaceType int

const (
	SURFACE_TYPE_IMAGE SurfaceType = iota
	SURFACE_TYPE_PDF
	SURFACE_TYPE_PS
	SURFACE_TYPE_XLIB
	SURFACE_TYPE_XCB
	SURFACE_TYPE_GLITZ
	SURFACE_TYPE_QUARTZ
	SURFACE_TYPE_WIN32
	SURFACE_TYPE_BEOS
	SURFACE_TYPE_DIRECTFB
	SURFACE_TYPE_SVG
	SURFACE_TYPE_OS2
	SURFACE_TYPE_WIN32_PRINTING
	SURFACE_TYPE_QUARTZ_IMAGE
	SURFACE_TYPE_SCRIPT
	SURFACE_TYPE_QT
	SURFACE_TYPE_RECORDING
	SURFACE_TYPE_VG
	SURFACE_TYPE_GL
	SURFACE_TYPE_DRM
	SURFACE_TYPE_TEE
	SURFACE_TYPE_XML
	SURFACE_TYPE_SKIA
	SURFACE_TYPE_SUBSURFACE
)

type Format int

func (self Format) StrideForWidth(width int) int {
	return int(C.cairo_format_stride_for_width(C.cairo_format_t(self), C.int(width)))
}

// cairo_format_t values
const (
	FORMAT_INVALID   Format = -1
	FORMAT_ARGB32    Format = 0
	FORMAT_RGB24     Format = 1
	FORMAT_A8        Format = 2
	FORMAT_A1        Format = 3
	FORMAT_RGB16_565 Format = 4
	FORMAT_RGB30     Format = 5
)

// cairo_pattern_type_t
type PatternType int

const (
	PATTERN_TYPE_SOLID PatternType = iota
	PATTERN_TYPE_SURFACE
	PATTERN_TYPE_LINEAR
	PATTERN_TYPE_RADIAL
)

// cairo_extend_t
type Extent int

const (
	EXTEND_NONE Extent = iota
	EXTEND_REPEAT
	EXTEND_REFLECT
	EXTEND_PAD
)

// cairo_filter_t
type Filter int

const (
	CAIRO_FILTER_FAST Filter = iota
	CAIRO_FILTER_GOOD
	CAIRO_FILTER_BEST
	CAIRO_FILTER_NEAREST
	CAIRO_FILTER_BILINEAR
	CAIRO_FILTER_GAUSSIAN
)

type RegionOverlap int

const (
	REGION_OVERLAP_IN RegionOverlap = iota
	REGION_OVERLAP_OUT
	REGION_OVERLAP_PART
)

type DeviceType int

const (
	DEVICE_TYPE_DRM DeviceType = iota
	DEVICE_TYPE_GL
	DEVICE_TYPE_SCRIPT
	DEVICE_TYPE_XCB
	DEVICE_TYPE_XLIB
	DEVICE_TYPE_XML
)

const (
	MIME_TYPE_JPEG = "image/jpeg"
	MIME_TYPE_PNG  = "image/png"
	MIME_TYPE_JP2  = "image/jp2"
	MIME_TYPE_URI  = "text/x-uri"
)

type PDFVersion int

func (self PDFVersion) String() string {
	return C.GoString(C.cairo_pdf_version_to_string(C.cairo_pdf_version_t(self)))
}

const (
	PDF_VERSION_1_4 PDFVersion = iota
	PDF_VERSION_1_5
)

type PSLevel int

func (self PSLevel) String() string {
	return C.GoString(C.cairo_ps_level_to_string(C.cairo_ps_level_t(self)))
}

const (
	PS_LEVEL_2 PSLevel = iota
	PS_LEVEL_3
)

type SVGVersion int

func (self SVGVersion) String() string {
	return C.GoString(C.cairo_svg_version_to_string(C.cairo_svg_version_t(self)))
}

const (
	SVG_VERSION_1_1 SVGVersion = iota
	SVG_VERSION_1_2
)

type Pattern struct {
	pattern *C.cairo_pattern_t
}

type Linear struct {
	X0, Y0 float64
	X1, Y1 float64
}

type Radial struct {
	CX0, CY0, Radius0 float64
	CX1, CY1, Radius1 float64
}

type Rectangle struct {
	X, Y          float64
	Width, Height float64
}

type TextCluster struct {
	// todo
}

type TextExtents struct {
	Xbearing float64
	Ybearing float64
	Width    float64
	Height   float64
	Xadvance float64
	Yadvance float64
}

type FontExtents struct {
	Ascent      float64
	Descent     float64
	Height      float64
	MaxXadvance float64
	MaxYadvance float64
}

type FontFace struct {
	face    *C.cairo_font_face_t
	ft_face *C.FT_Face
}

type FontOptions struct {
	// todo
}

type ScaledFont struct {
	// todo
}

type Glyph struct {
	// todo
}

type Device struct {
}

func Version() int {
	return int(C.cairo_version())
}

func VersionString() string {
	return C.GoString(C.cairo_version_string())
}
