// package cairo contains constants for use in cairo
package cairo

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

// Device cairo_device_t
type Device struct {
}

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
