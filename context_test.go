// Package cairo wraps the c cairographics library.
package cairo

import (
	"math"
	"testing"
)

func createContext() (*Surface, *Context) {
	surface := NewSurface(400, 400)
	context := NewContext(surface)
	return surface, context
}

func getPixel(data []byte, x, y, w int) (byte, byte, byte, byte) {
	index := (y*400 + x) * 4
	blue := data[index]
	green := data[index+1]
	red := data[index+2]
	alpha := data[index+3]
	return red, green, blue, alpha
}

func getData(surface *Surface, t *testing.T) []byte {
	t.Helper()
	data, err := surface.GetData()
	if err != nil {
		t.Errorf("Unable to get image data. Error: %s\n", err)
	}
	return data
}

func TestContextCreation(t *testing.T) {
	surface := NewSurface(400, 400)
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
	_, context := createContext()

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

func TestSaveAndRestore(t *testing.T) {
	surface, context := createContext()

	// translate and draw red at 0, 0, 10, 10
	context.Save()
	context.Translate(100, 100)
	context.SetSourceRGB(1, 0, 0)
	context.Rectangle(0, 0, 10, 10)
	context.Fill()

	// restore and draw blue at 0, 0, 10, 10
	context.Restore()
	context.SetSourceRGB(0, 0, 1)
	context.Rectangle(0, 0, 10, 10)
	context.Fill()

	data := getData(surface, t)

	// find red at translated position
	x := 105
	y := 105
	r, _, _, _ := getPixel(data, x, y, 400)
	if r != 255 {
		t.Errorf("Expected pixel 105, 105 to have red of 255, got %d\n", r)
	}

	// find blue at untranslated position
	x = 5
	y = 5
	_, _, b, _ := getPixel(data, x, y, 400)
	if r != 255 {
		t.Errorf("Expected pixel 105, 105 to have red of 255, got %d\n", r)
	}
	if b != 255 {
		t.Errorf("Expected pixel 5, 5 to have blue of 255, got %d\n", b)
	}
}

func TestOperator(t *testing.T) {
	surface, context := createContext()

	context.SetSourceRGB(1, 0, 0)
	context.Rectangle(100, 100, 100, 100)
	context.Fill()

	// only testing one operator here to make sure that setting operators works in general.
	// testing this specific method, not all of cairo's operators.
	// OperatorDestOver will draw blue square under red square.
	context.SetOperator(OperatorDestOver)
	context.SetSourceRGB(0, 0, 1)
	context.Rectangle(150, 150, 100, 100)
	context.Fill()

	data := getData(surface, t)

	// though the blue square was drawn here, it will be under the red square
	// so this pixel should be red
	x := 175
	y := 175
	r, _, b, _ := getPixel(data, x, y, 400)
	if r != 255 {
		t.Errorf("Expected pixel 175, 175 to have red of 255, got %d\n", r)
	}
	if b != 0 {
		t.Errorf("Expected pixel 175, 175 to have blue of 0, got %d\n", b)
	}

	// this is outside of the red square, so should be blue
	x = 225
	y = 225
	r, _, b, _ = getPixel(data, x, y, 400)
	if r != 0 {
		t.Errorf("Expected pixel 225, 225 to have red of 0, got %d\n", r)
	}
	if b != 255 {
		t.Errorf("Expected pixel 225, 225 to have blue of 255, got %d\n", b)
	}
}

func TestSetSourceRGB(t *testing.T) {
	surface, context := createContext()
	context.SetSourceRGB(1, 0, 0)
	context.Rectangle(0, 0, 10, 10)
	context.Fill()

	context.SetSourceRGB(0, 1, 0)
	context.Rectangle(10, 10, 10, 10)
	context.Fill()

	context.SetSourceRGB(0, 0, 1)
	context.Rectangle(20, 20, 10, 10)
	context.Fill()

	data := getData(surface, t)

	// though the blue square was drawn here, it will be under the red square
	// so this pixel should be red
	x := 5
	y := 5
	r, g, b, _ := getPixel(data, x, y, 400)
	if r != 255 {
		t.Errorf("Expected pixel %d, %d to have red of 255, got %d\n", x, y, r)
	}
	if g != 0 {
		t.Errorf("Expected pixel %d, %d to have red of 0, got %d\n", x, y, g)
	}
	if b != 0 {
		t.Errorf("Expected pixel %d, %d to have blue of 0, got %d\n", x, y, b)
	}

	x = 15
	y = 15
	red, green, blue, _ := getPixel(data, x, y, 400)
	if red != 0 {
		t.Errorf("Expected pixel %d, %d to have red of 0, got %d\n", x, y, red)
	}
	if green != 255 {
		t.Errorf("Expected pixel %d, %d to have red of 255, got %d\n", x, y, green)
	}
	if blue != 0 {
		t.Errorf("Expected pixel %d, %d to have blue of 0, got %d\n", x, y, blue)
	}

	x = 25
	y = 25
	red, green, blue, _ = getPixel(data, x, y, 400)
	if red != 0 {
		t.Errorf("Expected pixel %d, %d to have red of 0, got %d\n", x, y, red)
	}
	if green != 0 {
		t.Errorf("Expected pixel %d, %d to have red of 0, got %d\n", x, y, green)
	}
	if blue != 255 {
		t.Errorf("Expected pixel %d, %d to have blue of 255, got %d\n", x, y, blue)
	}

}

func TestSetSourceRGBA(t *testing.T) {
	surface, context := createContext()
	context.SetSourceRGBA(1, 0, 0, 0)
	context.Rectangle(0, 0, 10, 10)
	context.Fill()

	context.SetSourceRGBA(1, 0, 0, 0.5)
	context.Rectangle(10, 10, 10, 10)
	context.Fill()

	context.SetSourceRGBA(1, 0, 0, 1)
	context.Rectangle(20, 20, 10, 10)
	context.Fill()

	data := getData(surface, t)

	// only checking alpha here. tested rgb above.
	// and rgb values with less than 1.0 alpha will be premultiplied, so not what was set.
	x := 5
	y := 5
	_, _, _, alpha := getPixel(data, x, y, 400)
	if alpha != 0 {
		t.Errorf("Expected pixel %d, %d to have alpha of 0, got %d\n", x, y, alpha)
	}

	x = 15
	y = 15
	_, _, _, alpha = getPixel(data, x, y, 400)
	if alpha != 128 {
		t.Errorf("Expected pixel %d, %d to have alpha of 128, got %d\n", x, y, alpha)
	}

	x = 25
	y = 25
	_, _, _, alpha = getPixel(data, x, y, 400)
	if alpha != 255 {
		t.Errorf("Expected pixel %d, %d to have alpha of 255, got %d\n", x, y, alpha)
	}
}

func TestLineWidth(t *testing.T) {
	surface, context := createContext()

	context.SetLineWidth(10)
	w := context.GetLineWidth()

	if w != 10 {
		t.Errorf("Expected line width to be  %f, got %f\n", 10.0, w)
	}
	context.MoveTo(10, 10)
	context.LineTo(20, 10)
	context.Stroke()

	context.SetLineWidth(1)
	w = context.GetLineWidth()

	if w != 1 {
		t.Errorf("Expected line width to be  %f, got %f\n", 1.0, w)
	}
	context.MoveTo(10, 30)
	context.LineTo(20, 30)
	context.Stroke()

	data := getData(surface, t)

	x := 15
	y := 13
	_, _, _, alpha := getPixel(data, x, y, 400)
	if alpha == 0 {
		t.Errorf("Expected pixel %d, %d to have alpha greater than 0, got %d\n", x, y, alpha)
	}

	x = 15
	y = 33
	_, _, _, alpha = getPixel(data, x, y, 400)
	if alpha != 0 {
		t.Errorf("Expected pixel %d, %d to have alpha of 0, got %d\n", x, y, alpha)
	}
}

func TestScale(t *testing.T) {
	surface, context := createContext()
	context.Scale(10, 10)
	context.Rectangle(1, 1, 100, 100)
	context.Fill()

	data := getData(surface, t)

	// rect would cover this area, but scaling makes it not do so, so alpha should be zero.
	x := 5
	y := 5
	_, _, _, alpha := getPixel(data, x, y, 400)
	if alpha != 0 {
		t.Errorf("Expected pixel %d, %d to have alpha of 0, got %d\n", x, y, alpha)
	}

	// This would be outside the rect, but scaling makes it inside.
	x = 105
	y = 105
	_, _, _, alpha = getPixel(data, x, y, 400)
	if alpha == 0 {
		t.Errorf("Expected pixel %d, %d to have alpha greater than 0, got %d\n", x, y, alpha)
	}
}

func TestRotate(t *testing.T) {
	surface, context := createContext()
	context.Translate(50, 50)
	context.Rotate(math.Pi / 4)
	context.Rectangle(-50, -50, 100, 100)
	context.Fill()

	data := getData(surface, t)

	// rect would cover this area, but rotation makes it not do so, so alpha should be zero.
	x := 5
	y := 5
	_, _, _, alpha := getPixel(data, x, y, 400)
	if alpha != 0 {
		t.Errorf("Expected pixel %d, %d to have alpha of 0, got %d\n", x, y, alpha)
	}

	// This would be outside the rect, but rotation makes it inside.
	x = 105
	y = 50
	_, _, _, alpha = getPixel(data, x, y, 400)
	if alpha == 0 {
		t.Errorf("Expected pixel %d, %d to have alpha greater than 0, got %d\n", x, y, alpha)
	}
}

func TestCurveTo(t *testing.T) {
	surface, context := createContext()
	context.SetLineWidth(20)
	context.MoveTo(0, 400)
	context.CurveTo(0, 0, 400, 0, 400, 400)
	context.Stroke()

	data := getData(surface, t)

	// on the curve
	x := 200
	y := 100
	_, _, _, alpha := getPixel(data, x, y, 400)
	if alpha == 0 {
		t.Errorf("Expected pixel %d, %d to have alpha greater than 0, got %d\n", x, y, alpha)
	}

	// above the curve
	x = 200
	y = 50
	_, _, _, alpha = getPixel(data, x, y, 400)
	if alpha != 0 {
		t.Errorf("Expected pixel %d, %d to have alpha 0, got %d\n", x, y, alpha)
	}

	// below the curve
	x = 200
	y = 150
	_, _, _, alpha = getPixel(data, x, y, 400)
	if alpha != 0 {
		t.Errorf("Expected pixel %d, %d to have alpha 0, got %d\n", x, y, alpha)
	}
}

func TestArcs(t *testing.T) {
	surface, context := createContext()
	context.SetLineWidth(20)
	context.Arc(200, 200, 100, 0, math.Pi)
	context.Stroke()

	data := getData(surface, t)

	// on the curve
	x := 200
	y := 300
	_, _, _, alpha := getPixel(data, x, y, 400)
	if alpha == 0 {
		t.Errorf("Expected pixel %d, %d to have alpha greater than 0, got %d\n", x, y, alpha)
	}

	// not on the curve
	x = 200
	y = 100
	_, _, _, alpha = getPixel(data, x, y, 400)
	if alpha != 0 {
		t.Errorf("Expected pixel %d, %d to have alpha 0, got %d\n", x, y, alpha)
	}

	context.ArcNegative(200, 200, 100, 0, math.Pi)
	context.Stroke()

	data = getData(surface, t)
	// now it is on the curve
	x = 200
	y = 100
	_, _, _, alpha = getPixel(data, x, y, 400)
	if alpha == 0 {
		t.Errorf("Expected pixel %d, %d to have alpha greater than 0, got %d\n", x, y, alpha)
	}

}

func TestRelativeDrawing(t *testing.T) {
	_, context := createContext()
	context.MoveTo(100, 100)
	context.RelMoveTo(50, 50)
	x, y := context.GetCurrentPoint()
	if x != 150 && y != 150 {
		t.Errorf("Expected current point to be %f, %f, got  %f, %f\n", 150.0, 150.0, x, y)
	}

	context.RelLineTo(-50, 50)
	x, y = context.GetCurrentPoint()
	if x != 100 && y != 200 {
		t.Errorf("Expected current point to be %f, %f, got  %f, %f\n", 100.0, 200.0, x, y)
	}

	context.RelCurveTo(100, 100, 150, 100, 200, 0)
	x, y = context.GetCurrentPoint()
	if x != 300 && y != 200 {
		t.Errorf("Expected current point to be %f, %f, got  %f, %f\n", 300.0, 200.0, x, y)
	}
}

func TestClosePath(t *testing.T) {
	surface, context := createContext()
	context.SetLineWidth(20)
	context.MoveTo(100, 100)
	context.LineTo(200, 200)
	context.LineTo(100, 300)
	context.ClosePath()
	context.Stroke()

	data := getData(surface, t)

	// no line was drawn here, but closing the path will draw it.
	x := 100
	y := 200
	_, _, _, alpha := getPixel(data, x, y, 400)
	if alpha == 0 {
		t.Errorf("Expected pixel %d, %d to have alpha greater than 0, got %d\n", x, y, alpha)
	}
}

func TestStrokePreserve(t *testing.T) {
	surface, context := createContext()
	context.Rectangle(100, 100, 100, 100)
	context.Stroke()
	context.Fill()

	data := getData(surface, t)

	// rect will not be filled here.
	x := 150
	y := 150
	_, _, _, alpha := getPixel(data, x, y, 400)
	if alpha != 0 {
		t.Errorf("Expected pixel %d, %d to have alpha 0, got %d\n", x, y, alpha)
	}

	context.Rectangle(100, 100, 100, 100)
	context.StrokePreserve()
	context.Fill()

	data = getData(surface, t)

	// now rect will be filled
	x = 150
	y = 150
	_, _, _, alpha = getPixel(data, x, y, 400)
	if alpha == 0 {
		t.Errorf("Expected pixel %d, %d to have alpha greater than 0, got %d\n", x, y, alpha)
	}
}

func TestFillPreserve(t *testing.T) {
	surface, context := createContext()

	context.SetLineWidth(20)
	context.Rectangle(100, 100, 100, 100)
	context.Fill()
	context.Stroke()

	data := getData(surface, t)

	// rect will not be stroked here.
	x := 95
	y := 150
	_, _, _, alpha := getPixel(data, x, y, 400)
	if alpha != 0 {
		t.Errorf("Expected pixel %d, %d to have alpha 0, got %d\n", x, y, alpha)
	}

	context.Rectangle(100, 100, 100, 100)
	context.FillPreserve()
	context.Stroke()

	data = getData(surface, t)

	// now rect will be stroked
	x = 95
	y = 150
	_, _, _, alpha = getPixel(data, x, y, 400)
	if alpha == 0 {
		t.Errorf("Expected pixel %d, %d to have alpha greater than 0, got %d\n", x, y, alpha)
	}
}
