## go-cairo

### Go binding for the cairo graphics library

Based on Dethe Elza's version https://bitbucket.org/dethe/gocairo
but significantly extended and updated.

Go specific extensions:
* NewSurfaceFromImage(image.Image)
* Surface.GetData() []byte
* Surface.SetData([]byte)
* Surface.GetImage() image.Image
* Surface.SetImage(image.Image)

go-cairo also sports a sub package extimage with image.Image/draw.Image
implementations for 32 bit ARGB and 24 bit RGB color models.

Overview:
* http://go.pkgdoc.org/github.com/ungerik/go-cairo
* http://go.pkgdoc.org/github.com/ungerik/go-cairo/extimage

Missing features
* TextCluster
* FontExtents
* FontFace
* FontOptions
* ScaledFont
* Glyph

### Installation:

Install cairo:

For Debian and Debian derivatives including Ubuntu:

	sudo apt-get install libcairo2-dev

For Fedora:

	sudo yum install cairo-devel

For openSUSE:

	zypper install cairo-devel

For Arch Linux: (also [gcc](https://archlinux.org/packages/core/x86_64/gcc/) and [pkgconf](https://archlinux.org/packages/core/x86_64/pkgconf/) are required)
	
	sudo pacman -S cairo

For MacOS you may get errors finding cairo-pdf.h. To fix this run:

	xcode-select --install
  
Install go-cairo and run go-cairo-example:

	go get github.com/ungerik/go-cairo
	go install github.com/ungerik/go-cairo/go-cairo-example && go-cairo-example


Copyrights: See LICENSE file


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/ungerik/go-cairo/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

