// Package cairo wraps the c cairographics library.
package cairo

import (
	"errors"
	"os"
	"testing"
)

func TestSurfaceCreation(t *testing.T) {
	surface := NewSurface(FormatARGB32, 800, 800)

	status := surface.GetStatus()
	if status != StatusSuccess {
		t.Errorf("Expected status %v, got %v\n", StatusSuccess, status)
	}

	format := surface.GetFormat()
	if format != FormatARGB32 {
		t.Errorf("Expected format %v, got %v\n", FormatARGB32, format)
	}

	content := surface.GetContent()
	if content != ContentColorAlpha {
		t.Errorf("Expected content type %v, got %v\n", ContentColorAlpha, content)
	}

	expected := 800 * 4
	stride := surface.GetStride()
	if stride != expected {
		t.Errorf("Expected stride %d, got %d\n", expected, stride)
	}

	width, height := surface.GetWidth(), surface.GetHeight()
	if width != 800 {
		t.Errorf("Expected width %d, got %d\n", 800, width)
	}
	if height != 800 {
		t.Errorf("Expected height %d, got %d\n", 800, height)
	}
}

func TestSurfaceFromPng(t *testing.T) {
	surface, err := NewSurfaceFromPNG("testdata/gopher.png")
	if err != nil {
		t.Errorf("Unable to create png. Error: %v\n", err)
	}

	status := surface.GetStatus()
	if status != StatusSuccess {
		t.Errorf("Expected status %v, got %v\n", StatusSuccess, status)
	}

	format := surface.GetFormat()
	if format != FormatRGB24 {
		t.Errorf("Expected format %v, got %v\n", FormatRGB24, format)
	}

	content := surface.GetContent()
	if content != ContentColor {
		t.Errorf("Expected content type %v, got %v\n", ContentColor, content)
	}

	expected := 666 * 4
	stride := surface.GetStride()
	if stride != expected {
		t.Errorf("Expected stride %d, got %d\n", expected, stride)
	}

	width, height := surface.GetWidth(), surface.GetHeight()
	if width != 666 {
		t.Errorf("Expected width %d, got %d\n", 666, width)
	}
	if height != 915 {
		t.Errorf("Expected height %d, got %d\n", 915, height)
	}

}

func TestFinish(t *testing.T) {
	surface := NewSurface(FormatARGB32, 800, 800)
	context := NewContext(surface)
	context.Rectangle(0, 0, 100, 100)
	context.Stroke()
	status := surface.GetStatus()
	if status != StatusSuccess {
		t.Errorf("Expected status %q, got %q\n", StatusSuccess, status)
	}

	surface.Finish()
	context.Rectangle(0, 0, 100, 100)
	context.Stroke()
	status = surface.GetStatus()
	if status != StatusSurfaceFinished {
		t.Errorf("Expected status %q, got %q\n", StatusSurfaceFinished, status)
	}
}

func TestWritePng(t *testing.T) {
	path := "testdata/temp.png"
	surface := NewSurface(FormatARGB32, 800, 800)
	err := surface.WriteToPNG(path)
	if err != nil {
		t.Errorf("Unable to save image. Error: %s\n", err)
	}
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		t.Errorf("File was not saved.")
	}
	os.Remove(path)
}

func TestGetData(t *testing.T) {
	surface := NewSurface(FormatARGB32, 10, 10)
	context := NewContext(surface)
	context.SetSourceRGB(1, 0, 0)
	context.Rectangle(0, 0, 10, 10)
	context.Fill()

	data, err := surface.GetData()
	if err != nil {
		t.Errorf("Unable to get image data. Error: %s\n", err)
	}
	expected := 10 * 10 * 4
	dataLength := len(data)
	if dataLength != expected {
		t.Errorf("Expected data length of %d, got %d\n", expected, dataLength)
	}

	// blue
	if data[0] != 0 {
		t.Errorf("Expected data[0] to be %d, got %d\n", 0, data[0])
	}
	// green
	if data[1] != 0 {
		t.Errorf("Expected data[0] to be %d, got %d\n", 0, data[1])
	}
	// red
	if data[2] != 255 {
		t.Errorf("Expected data[0] to be %d, got %d\n", 255, data[2])
	}
	// alpha
	if data[3] != 255 {
		t.Errorf("Expected data[0] to be %d, got %d\n", 255, data[3])
	}
}

func TestSetData(t *testing.T) {
	surface := NewSurface(FormatARGB32, 10, 10)

	data := make([]byte, 10*10*4)
	data[99] = 199
	surface.SetData(data)

	gotData, err := surface.GetData()
	if err != nil {
		t.Errorf("Unable to get image data. Error: %s\n", err)
	}
	if gotData[99] != 199 {
		t.Errorf("Expected data[99] to be %d, got %d\n", 199, data[99])
	}
}
