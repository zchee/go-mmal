// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

const (
	// VersionMajor Major version number.
	// This changes when the API breaks in a way which is not backward compatible.
	VersionMajor = 0
	// VersionMinor Minor version number.
	// This changes each time the API is extended in a way which is still source and
	// binary compatible.
	VersionMinor = 0
	// Version version number.
	Version = VersionMajor<<16 | VersionMinor
)

// VersionToMajor returns major mmal version from a.
func VersionToMajor(a int) int {
	return a >> 16
}

// VersionToMinor returns minor mmal version from a.
func VersionToMinor(a int) int {
	return a & 0xFFFF
}
