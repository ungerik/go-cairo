## go-cairo

### Go binding for the cairo graphics library

Very significant fork from Erik Unger's version https://github.com/ungerik/go-cairo
* Added a a lot of missing methods.
* Broke up surface and context to separate entities in separate files.
* Broke down cairo.go into a lot of individual const files.
* General naming best practices: receiver names, var and const names, function names, etc.
* Did a lot of documentation.
* Removed a bunch of dead or stubbed out code that wasn't being used.
* Moved from panics and cairo statuses to Go errors.
* Began writing tests.

### Installation:

1. Install cairo:

  * For Debian and Debian derivatives including Ubuntu:

    `sudo apt-get install libcairo2-dev`

  * For Fedora:

    `sudo yum install cairo-devel`

  * For openSUSE:

    `zypper install cairo-devel`
          
  * For MacOS HomeBrew:

    `brew install cairo`

  * For MacOS MacPorts:

    `sudo port install cairo`

  * For Windows, install gtk:

    `http://www.gtk.org/download/`

2. Install go-cairo

  `go get github.com/bit101/go-cairo`

3. Check examples in `go-cairo-examples`

Copyrights: See LICENSE file
