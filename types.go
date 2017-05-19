// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal_types.h>
*/
import "C"

type Status C.MMAL_STATUS_T

const (
	Success Status = iota
	ErrNoMem
	ErrNoSpc
	ErrInVal
	ErrNoSys
	ErrNoEnt
	ErrNXIO
	ErrIO
	ErrSPipe
	ErrCorrupt
	ErrNotReady
	ErrConfig
	ErrIsConn
	ErrNotConn
	ErrAgain
	ErrFault
	ErrMax = 0x7FFFFFFF
)

type Rect struct {
	c C.MMAL_RECT_T
}

func (r Rect) X() int32 {
	return int32(r.c.x)
}

func (r Rect) Y() int32 {
	return int32(r.c.y)
}

func (r Rect) Width() int32 {
	return int32(r.c.width)
}

func (r Rect) Height() int32 {
	return int32(r.c.height)
}

type Rational struct {
	c C.MMAL_RATIONAL_T
}

func (r Rational) Num() int32 {
	return int32(r.c.num)
}

func (r Rational) Den() int32 {
	return int32(r.c.den)
}

const TimeUnknown = 1 << 63

type FourCCT uint32
