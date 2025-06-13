[![Build Status](https://travis-ci.com/evertrus/clipboard.svg?branch=master)](https://travis-ci.com/evertrus/clipboard)

[![GoDoc](https://godoc.org/github.com/evertrus/clipboard?status.svg)](http://godoc.org/github.com/evertrus/clipboard)

# Clipboard for Go

Provide copying and pasting to the Clipboard for Go.

Build:

    $ go get github.com/evertrus/clipboard

Platforms:

* OSX
* Windows 7 (probably work on other Windows)
* Linux, Unix (requires 'xclip' or 'xsel' command to be installed)


Document: 

* http://godoc.org/github.com/evertrus/clipboard

Notes:

* Text string only
* UTF-8 text encoding only (no conversion)

TODO:

* Clipboard watcher(?)

## Commands:

paste shell command:

    $ go get github.com/evertrus/clipboard/cmd/gopaste
    $ # example:
    $ gopaste > document.txt

copy shell command:

    $ go get github.com/evertrus/clipboard/cmd/gocopy
    $ # example:
    $ cat document.txt | gocopy

## Credits
This is a fork of [atotto/clipboard](https://github.com/atotto/clipboard) that enables support for AIX compilation

