// package cairo contains constants for use in cairo
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// MimeType constants
const (
	MimeTypeJPEG = "image/jpeg"
	MimeTypePNG  = "image/png"
	MimeTypeJP2  = "image/jp2"
	MimeTypeURI  = "text/x-uri"
)
