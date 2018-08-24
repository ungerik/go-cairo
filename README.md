## go-cairo

### Go binding for the cairo graphics library

Forked from Erik Unger's version https://github.com/ungerik/go-cairo
* added missing methods, including support for radial gradients and mesh patterns
* began to flesh out matrix and transform methods more completely
* cleaned up code to pass the Go linter and Go best practices in general
* began to flesh out documentation of all methods
* removed `extimage` package, which, while useful, transcends cairo itself

Missing features
* TextCluster
* FontExtents
* FontFace
* FontOptions
* ScaledFont
* Glyph
* GetDevice(s)

Future plans
* Possibly separate surface and context to be more in line with the cairo structure
* Finish documenting methods

### Installation:

Install cairo:

* For Debian and Debian derivatives including Ubuntu:

	sudo apt-get install libcairo2-dev

* For Fedora:

	sudo yum install cairo-devel

* For openSUSE:

	zypper install cairo-devel
	
* For MacOS HomeBrew:

	brew install cairo

* For MacOS MacPorts:

	sudo port install cairo

* For Windows, install gtk:

	http://www.gtk.org/download/

Install go-cairo

	go get github.com/bit101/go-cairo

Check examples at `go-cairo-examples`

Copyrights: See LICENSE file
