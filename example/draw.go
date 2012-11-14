package main

// Simple test for cairo package

import "github.com/ungerik/go-cairo"

/*
Go port of this C code:

cairo_surface_t *surface =
      cairo_image_surface_create (CAIRO_FORMAT_ARGB32, 240, 80);
  cairo_t *cr =
      cairo_create (surface);

  cairo_select_font_face (cr, "serif", CAIRO_FONT_SLANT_NORMAL, CAIRO_FONT_WEIGHT_BOLD);
  cairo_set_font_size (cr, 32.0);
  cairo_set_source_rgb (cr, 0.0, 0.0, 1.0);
  cairo_move_to (cr, 10.0, 50.0);
  cairo_show_text (cr, "Hello, world");

  cairo_destroy (cr);
  cairo_surface_write_to_png (surface, "hello.png");
  cairo_surface_destroy (surface);

*/

func main() {
	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, 240, 80)
	surface.SelectFontFace("serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surface.SetFontSize(32.0)
	surface.SetSourceRGB(0.0, 0.0, 1.0)
	surface.MoveTo(10.0, 50.0)
	surface.ShowText("Hello World")
	surface.WriteToPNG("hello.png")
	surface.Finish()
}
