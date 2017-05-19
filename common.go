// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal_common.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func ToString(s interface{}) string {
	return fmt.Sprint(s)
}

func CountOf(x unsafe.Pointer) interface{} { return nil }

func FourCC(a, b, c, d rune) int {
	return int(a) | int(b)<<8 | int(c)<<16 | int(d)<<24
}

var Magic = FourCC('m', 'm', 'a', 'l')

type BoolT C.MMAL_BOOL_T

var (
	False = 0
	True  = 1
)

type CoreStatistics struct {
	BufferCount     uint32
	FirstBufferTime uint32
	LastBufferTime  uint32
	MaxDelay        uint32
}

type CorePortStatistics struct {
	RX CoreStatistics
	TX CoreStatistics
}

type Fixed1616 uint32
