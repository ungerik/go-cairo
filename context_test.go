// Package cairo wraps the c cairographics library.
package cairo

import (
	"testing"
)

func TestContextCreation(t *testing.T) {
	surface := NewSurface(FormatARGB32, 400, 400)
	context := NewContext(surface)

	status := context.GetStatus()
	if status != StatusSuccess {
		t.Errorf("Could not create context. Status: %q\n", status)
	}

	pngSurface, err := NewSurfaceFromPNG("testdata/gopher.png")
	if err != nil {
		t.Errorf("Unable to create png. Error: %v\n", err)
	}
	pngContext := NewContext(pngSurface)

	status = pngContext.GetStatus()
	if status != StatusSuccess {
		t.Errorf("Could not create context. Status: %q\n", status)
	}
}

func TestCurrentPoint(t *testing.T) {
	surface := NewSurface(FormatARGB32, 400, 400)
	context := NewContext(surface)

	hasPoint := context.HasCurrentPoint()
	if hasPoint {
		t.Errorf("New context should not have a current point\n")
	}
	x, y := context.GetCurrentPoint()
	if x != 0 && y != 0 {
		t.Errorf("Empty current point should be 0, 0. Got %f, %f\n", x, y)
	}

	context.MoveTo(100, 100)
	hasPoint = context.HasCurrentPoint()
	x, y = context.GetCurrentPoint()
	if !hasPoint {
		t.Errorf("Context should have a current point\n")
	}
	if x != 100 && y != 100 {
		t.Errorf("Empty current point should be 100, 100. Got %f, %f\n", x, y)
	}

	context.LineTo(200, 300)
	hasPoint = context.HasCurrentPoint()
	x, y = context.GetCurrentPoint()
	if !hasPoint {
		t.Errorf("Context should have a current point\n")
	}
	if x != 200 && y != 300 {
		t.Errorf("Empty current point should be 200, 300. Got %f, %f\n", x, y)
	}

}
