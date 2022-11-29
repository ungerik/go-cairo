// Package main contains an executable
package main

// Simple test for cairo package

import (
	"math"
	"math/rand"

	cairo "github.com/bit101/go-cairo"
)

func main() {
	text()
}

func text() {
	surface := cairo.NewSurface(240, 80)
	context := cairo.NewContext(surface)
	context.SetSourceRGB(1, 1, 1)
	context.Paint()
	context.SelectFontFace("Arial", cairo.FontSlantNormal, cairo.FontWeightBold)
	context.SetFontSize(32.0)
	context.SetSourceRGB(0.0, 0.0, 1.0)
	context.MoveTo(10.0, 50.0)
	context.ShowText("Hello World")
	surface.WriteToPNG("out.png")
	surface.Finish()
}

func shapes() {
	surface := cairo.NewSurface(600, 230)
	context := cairo.NewContext(surface)
	context.SetSourceRGB(1, 1, 1)
	context.Paint()
	context.SetSourceRGB(0, 0, 0)

	context.Rectangle(10, 10, 100, 100)
	context.Fill()

	context.Rectangle(120, 10, 100, 100)
	context.Stroke()

	context.Arc(280, 60, 50, 0, math.Pi*2)
	context.Fill()

	context.Arc(390, 60, 50, 0, math.Pi*2)
	context.Stroke()

	for i := 0; i < 50; i++ {
		context.MoveTo(450+rand.Float64()*100, 10+rand.Float64()*100)
		context.LineTo(450+rand.Float64()*100, 10+rand.Float64()*100)
		context.Stroke()
	}

	context.MoveTo(10, 120)
	context.CurveTo(590, 120, 10, 220, 590, 220)
	context.Stroke()

	surface.WriteToPNG("out.png")
	surface.Finish()
}

func colors() {
	surface := cairo.NewSurface(600, 600)
	context := cairo.NewContext(surface)
	for i := 0.0; i < 100; i++ {
		for j := 0.0; j < 100; j++ {
			dist := math.Hypot(i*6-300, j*6-300)
			red := i / 100
			green := math.Max(0, 1.0-dist/200)
			blue := j / 100
			context.SetSourceRGB(red, green, blue)
			context.Rectangle(i*6, j*6, 6, 6)
			context.Fill()
		}
	}

	surface.WriteToPNG("out.png")
	surface.Finish()
}

func gradients() {
	surface := cairo.NewSurface(600, 300)
	context := cairo.NewContext(surface)
	radialPattern := cairo.CreateRadialGradient(150, 150, 0, 150, 150, 150)
	radialPattern.AddColorStopRGB(0, 1, 0, 0)
	radialPattern.AddColorStopRGB(1, 0, 0, 1)
	context.SetSource(radialPattern)
	context.Rectangle(0, 0, 300, 300)
	context.Fill()

	linearPattern := cairo.CreateLinearGradient(300, 0, 600, 300)
	linearPattern.AddColorStopRGB(0, 1, 0, 0)
	linearPattern.AddColorStopRGB(1, 0, 0, 1)
	context.SetSource(linearPattern)
	context.Rectangle(300, 0, 300, 300)
	context.Fill()

	surface.WriteToPNG("out.png")
	surface.Finish()
}

func mesh() {
	surface := cairo.NewSurface(600, 600)
	context := cairo.NewContext(surface)
	pattern := cairo.CreateMesh()

	pattern.BeginPatch()
	pattern.MoveTo(100, 100)
	pattern.LineTo(500, 100)
	pattern.LineTo(500, 500)
	pattern.LineTo(100, 500)

	pattern.SetCornerColorRGB(0, 1, 0, 0)
	pattern.SetCornerColorRGB(1, 0, 1, 0)
	pattern.SetCornerColorRGB(2, 0, 0, 1)
	pattern.SetCornerColorRGB(3, 1, 1, 0)
	pattern.EndPatch()

	context.SetSource(pattern)
	context.Rectangle(0, 0, 600, 600)
	context.Fill()

	surface.WriteToPNG("out.png")
	surface.Finish()
}
