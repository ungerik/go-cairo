// package cairo contains constants for use in cairo
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

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
