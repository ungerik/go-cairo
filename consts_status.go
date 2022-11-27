// package cairo contains constants for use in cairo
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
