package main

// Simple test for cairo package

import (
	"math"
	"math/rand"

	"github.com/bit101/go-cairo"
)

func main() {
	colors()
}

func text() {
	surface := cairo.NewSurface(cairo.FormatARGB32, 240, 80)
	surface.SelectFontFace("serif", cairo.FontSlantNormal, cairo.FontWeightBold)
	surface.SetFontSize(32.0)
	surface.SetSourceRGB(0.0, 0.0, 1.0)
	surface.MoveTo(10.0, 50.0)
	surface.ShowText("Hello World")
	surface.WriteToPNG("out.png")
	surface.Finish()
}

func shapes() {
	surface := cairo.NewSurface(cairo.FormatARGB32, 600, 230)
	surface.SetSourceRGB(1, 1, 1)
	surface.Paint()
	surface.SetSourceRGB(0, 0, 0)

	surface.Rectangle(10, 10, 100, 100)
	surface.Fill()

	surface.Rectangle(120, 10, 100, 100)
	surface.Stroke()

	surface.Arc(280, 60, 50, 0, math.Pi*2)
	surface.Fill()

	surface.Arc(390, 60, 50, 0, math.Pi*2)
	surface.Stroke()

	for i := 0; i < 50; i++ {
		surface.MoveTo(450+rand.Float64()*100, 10+rand.Float64()*100)
		surface.LineTo(450+rand.Float64()*100, 10+rand.Float64()*100)
		surface.Stroke()
	}

	surface.MoveTo(10, 120)
	surface.CurveTo(590, 120, 10, 220, 590, 220)
	surface.Stroke()

	surface.WriteToPNG("out.png")
	surface.Finish()
}

func colors() {
	surface := cairo.NewSurface(cairo.FormatARGB32, 600, 600)
	for i := 0.0; i < 100; i++ {
		for j := 0.0; j < 100; j++ {
			dist := math.Hypot(i*6-300, j*6-300)
			red := i / 100
			green := math.Max(0, 1.0-dist/200)
			blue := j / 100
			surface.SetSourceRGB(red, green, blue)
			surface.Rectangle(i*6, j*6, 6, 6)
			surface.Fill()
		}
	}

	surface.WriteToPNG("out.png")
	surface.Finish()
}
