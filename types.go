// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal_types.h>
*/
import "C"
import (
	"fmt"
)

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

func (s Status) String() string {
	switch s {
	case Success:
		return "Success"
	case ErrNoMem:
		return "ErrNoMem"
	case ErrNoSpc:
		return "ErrNoSpc"
	case ErrInVal:
		return "ErrInVal"
	case ErrNoSys:
		return "ErrNoSys"
	case ErrNoEnt:
		return "ErrNoEnt"
	case ErrNXIO:
		return "ErrNXIO"
	case ErrIO:
		return "ErrIO"
	case ErrSPipe:
		return "ErrSPipe"
	case ErrCorrupt:
		return "ErrCorrupt"
	case ErrNotReady:
		return "ErrNotReady"
	case ErrConfig:
		return "ErrConfig"
	case ErrIsConn:
		return "ErrIsConn"
	case ErrNotConn:
		return "ErrNotConn"
	case ErrAgain:
		return "ErrAgain"
	case ErrFault:
		return "ErrFault"
	case ErrMax:
		return "ErrMax"
	}
	return fmt.Sprintf("Status(%d)", s)
}

func (s Status) Error() string {
	switch s {
	case Success:
		return "Success"
	case ErrNoMem:
		return "Out of memory"
	case ErrNoSpc:
		return "Out of resources (other than memory)"
	case ErrInVal:
		return "Argument is invalid"
	case ErrNoSys:
		return "Function not implemented"
	case ErrNoEnt:
		return "No such file or directory"
	case ErrNXIO:
		return "No such device or address"
	case ErrIO:
		return "I/O error"
	case ErrSPipe:
		return "Illegal seek"
	case ErrCorrupt:
		return "Data is corrupt"
	case ErrNotReady:
		return "Component is not ready"
	case ErrConfig:
		return "Component is not configured"
	case ErrIsConn:
		return "Port is already connected"
	case ErrNotConn:
		return "Port is disconnected"
	case ErrAgain:
		return "Resource temporarily unavailable"
	case ErrFault:
		return "Bad address"
	case ErrMax:
		return "Force to 32 bit"
	}
	return "unknown error"
}

type Rect struct {
	c C.MMAL_RECT_T
}

func NewRect(x, y, width, height int32) Rect {
	var rect C.MMAL_RECT_T
	rect.x = C.int32_t(x)
	rect.y = C.int32_t(y)
	rect.width = C.int32_t(width)
	rect.height = C.int32_t(height)
	return Rect{rect}
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
