## go-cairo

### Go binding for the cairo graphics library

Based on Dethe Elza's version https://bitbucket.org/dethe/gocairo
but significantly extended and updated.

Forked from Erik Unger's version https://github.com/ungerik/go-cairo
but extended and updated a bit more.

Go specific extensions:
* NewSurfaceFromImage(image.Image)
* Surface.GetData() []byte
* Surface.SetData([]byte)
* Surface.GetImage() image.Image
* Surface.SetImage(image.Image)

go-cairo also sports a sub package extimage with image.Image/draw.Image
implementations for 32 bit ARGB and 24 bit RGB color models.

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
	
For MacOS:
HomeBrew:

	brew install cairo

MacPorts:

	sudo port install cairo

Windows:
Install gtk:

	http://www.gtk.org/download/

Install go-cairo and run go-cairo-example:

	go get github.com/bit101/go-cairo

Copyrights: See LICENSE file
