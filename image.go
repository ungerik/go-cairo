package cairo

import (
	"image"
)

func (self *Surface) GetImage() image.Image {
	width := self.GetWidth()
	height := self.GetHeight()
	stride := self.GetStride()
	data := self.GetData()

	switch self.GetFormat() {
	case FORMAT_ARGB32:
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		// could be optimized:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				surfOffset := y*stride + x*4
				imgOffset := y*img.Stride + x*4
				img.Pix[imgOffset+0] = data[surfOffset+1]
				img.Pix[imgOffset+1] = data[surfOffset+2]
				img.Pix[imgOffset+2] = data[surfOffset+3]
				img.Pix[imgOffset+3] = data[surfOffset+0]
			}
		}
		return img

	case FORMAT_RGB24:
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		// could be optimized:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				surfOffset := y*stride + x*3
				imgOffset := y*img.Stride + x*4
				img.Pix[imgOffset+0] = data[surfOffset+0]
				img.Pix[imgOffset+1] = data[surfOffset+1]
				img.Pix[imgOffset+2] = data[surfOffset+2]
				img.Pix[imgOffset+3] = 255
			}
		}
		return img

	case FORMAT_A8:
		img := image.NewAlpha(image.Rect(0, 0, width, height))
		// could be optimized:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				img.Pix[y*img.Stride+x] = data[y*stride+x]
			}
		}
		return img

	case FORMAT_A1:
		panic("Unsuppored surface format cairo.FORMAT_A1")

	case FORMAT_RGB16_565:
		panic("Unsuppored surface format cairo.FORMAT_RGB16_565")

	case FORMAT_RGB30:
		panic("Unsuppored surface format cairo.FORMAT_RGB30")

	case FORMAT_INVALID:
		panic("Invalid surface format")
	}
	panic("Unknown surface format")
}

// SetImage set the data from an image.Image with identical size.
func (self *Surface) SetImage(img image.Image) {
	width := self.GetWidth()
	height := self.GetHeight()
	if width != img.Bounds().Dx() || height != img.Bounds().Dy() {
		panic("Image size different from cairo surface size")
	}
	stride := self.GetStride()
	data := self.GetData()

	switch self.GetFormat() {
	case FORMAT_ARGB32:
		// could be optimized:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				offset := y*stride + x*4
				r, g, b, a := img.At(x, y).RGBA()
				data[offset+0] = byte(a >> 8)
				data[offset+1] = byte(r >> 8)
				data[offset+2] = byte(g >> 8)
				data[offset+3] = byte(b >> 8)
			}
		}
		self.SetData(data)

	case FORMAT_RGB24:
		// could be optimized:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				offset := y*stride + x*3
				r, g, b, _ := img.At(x, y).RGBA()
				data[offset+0] = byte(r >> 8)
				data[offset+1] = byte(g >> 8)
				data[offset+2] = byte(b >> 8)
			}
		}
		self.SetData(data)

	case FORMAT_A8:
		// could be optimized:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				_, _, _, a := img.At(x, y).RGBA()
				data[y*stride+x] = byte(a >> 8)
			}
		}
		self.SetData(data)

	case FORMAT_A1:
		panic("Unsuppored surface format cairo.FORMAT_A1")

	case FORMAT_RGB16_565:
		panic("Unsuppored surface format cairo.FORMAT_RGB16_565")

	case FORMAT_RGB30:
		panic("Unsuppored surface format cairo.FORMAT_RGB30")

	case FORMAT_INVALID:
		panic("Invalid surface format")
	}
	panic("Unknown surface format")
}
