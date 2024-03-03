//go:build !goci
// +build !goci

package cairo

/*
#include <cairo/cairo.h>
#include <cairo/cairo-ft.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type Cairo_freetype struct {
	library C.FT_Library
}

// Initialize a new FreeType library object
func InitFreeType() (Cairo_freetype, error) {
	var library C.FT_Library
	err := C.FT_Init_FreeType(&library)
	if err == 0 {
		return Cairo_freetype{library: library}, nil
	}
	return Cairo_freetype{}, fmt.Errorf("FT_Init_FreeType error %v", err)
}

// Destroy a given FreeType library object and all of its children, including resources, drivers, faces, sizes, etc.
func (ft Cairo_freetype) DoneFreeType() error {
	if ft.library != nil {
		err := C.FT_Done_FreeType(ft.library)
		ft.library = nil
		if err != 0 {
			return fmt.Errorf("FT_Done_FreeType error %v", err)
		}
	}
	return nil
}

// Call FT_Open_Face to open a font by its pathname
//
// Example:
//
//	//  Initialize a new FreeType library object
//	ft, err := cairo.InitFreeType()
//	if err != nil {
//	    log.Fatalln(err)
//	}
//	// Open font by its pathname
//	myfont, err := ft.FtNewFace("MyFont.ttf")
//	if err != nil {
//	    log.Fatalln(err)
//	}
//	// Create a surface:
//	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, 200, 200)
//	// Use the font:
//	surface.SetFontFace(myfont)
func (ft Cairo_freetype) FtNewFace(filename string) (*FontFace, error) {
	var face C.FT_Face
	cs := C.CString(filename)
	err := C.FT_New_Face(ft.library, cs, 0, &face)
	C.free(unsafe.Pointer(cs))
	if err != 0 {
		return nil, fmt.Errorf("FT_New_Face error %v", err)
	}

	return &FontFace{
		face:    C.cairo_ft_font_face_create_for_ft_face(face, 0),
		ft_face: &face,
	}, nil
}

// Call FT_Open_Face to open a font that has been loaded into memory
//
// This is similar to FtNewFace, but loads the font from memory instead
// of from file
func (ft Cairo_freetype) FtNewMemoryFace(data []byte) (*FontFace, error) {
	var face C.FT_Face
	err := C.FT_New_Memory_Face(ft.library, (*C.uchar)(&data[0]), C.long(len(data)), 0, &face)
	if err != 0 {
		return nil, fmt.Errorf("FT_New_Memory_Face error %v", err)
	}

	return &FontFace{
		face:    C.cairo_ft_font_face_create_for_ft_face(face, 0),
		ft_face: &face,
	}, nil
}

// Discard a given face object, as well as all of its child slots and sizes
func (ff *FontFace) FtDoneFace() error {
	if ff.face != nil {
		C.cairo_font_face_destroy(ff.face)
		ff.face = nil
	}
	if ff.ft_face != nil {
		err := C.FT_Done_Face(*ff.ft_face)
		ff.ft_face = nil
		if err != 0 {
			return fmt.Errorf("FT_Done_Face error %v", err)
		}
	}
	return nil
}
