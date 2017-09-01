package main

// VERSION set with Makefile using linker flag, must be uninitialized
var VERSION string

func init() {
	// provide a default version string if app is built without makefile
	if VERSION == "" {
		VERSION = "version-manually-built"
	}
}
