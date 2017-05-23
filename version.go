// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
*/
import "C"

const (
	// VersionMajor libmmal major version number.
	// This changes when the API breaks in a way which is not backward compatible.
	VersionMajor = C.MMAL_VERSION_MAJOR
	// VersionMinor libmmal minor version number.
	// This changes each time the API is extended in a way which is still source and binary compatible.
	VersionMinor = C.MMAL_VERSION_MINOR
	// Version libmmal version number.
	Version = C.MMAL_VERSION
)

// VersionToMajor returns major mmal version from a.
func VersionToMajor(a int) int {
	return a >> 16
}

// VersionToMinor returns minor mmal version from a.
func VersionToMinor(a int) int {
	return a & 0xFFFF
}
