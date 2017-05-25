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
)

func ToString(s interface{}) string {
	return fmt.Sprint(s)
}

// TODO(zchee): implements
// func CountOf(x unsafe.Pointer) interface{} { return nil }

func FourCC(a, b, c, d rune) int {
	return int(a) | int(b)<<8 | int(c)<<16 | int(d)<<24
}

var Magic = FourCC('m', 'm', 'a', 'l')

type BoolT C.MMAL_BOOL_T

var (
	False = 0
	True  = 1
)

type CoreStatisticsType struct {
	c C.MMAL_CORE_STATISTICS_T
}

func (c CoreStatisticsType) BufferCount() uint32 {
	return uint32(c.c.buffer_count)
}

func (c CoreStatisticsType) FirstBufferTime() uint32 {
	return uint32(c.c.first_buffer_time)
}

func (c CoreStatisticsType) LastBufferTime() uint32 {
	return uint32(c.c.last_buffer_time)
}

func (c CoreStatisticsType) MaxDelay() uint32 {
	return uint32(c.c.max_delay)
}

type CorePortStatisticsType struct {
	c C.MMAL_CORE_PORT_STATISTICS_T
}

func (c CorePortStatisticsType) RX() CoreStatisticsType {
	return CoreStatisticsType{c.c.rx}
}

func (c CorePortStatisticsType) TX() CoreStatisticsType {
	return CoreStatisticsType{c.c.tx}
}

type Fixed1616 uint32
