// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal_parameters_clock.h>
*/
import "C"

const (
	ParameterClockReference        ParameterGroupClock = (4 << 16) + iota // Takes a MMAL_PARAMETER_BOOLEAN_T
	ParameterClockActive                                                  // Takes a MMAL_PARAMETER_BOOLEAN_T
	ParameterClockScale                                                   // Takes a MMAL_PARAMETER_RATIONAL_T
	ParameterClockTime                                                    // Takes a MMAL_PARAMETER_INT64_T
	ParameterClockUpdateThreshold                                         // Takes a MMAL_PARAMETER_CLOCK_UPDATE_THRESHOLD_T
	ParameterClockDiscontThreshold                                        // Takes a MMAL_PARAMETER_CLOCK_DISCONT_THRESHOLD_T
	ParameterClockRequestThreshold                                        // Takes a MMAL_PARAMETER_CLOCK_REQUEST_THRESHOLD_T
	ParameterClockEnableBufferInfo                                        // Takes a MMAL_PARAMETER_BOOLEAN_T
	ParameterClockFrameRate                                               // Takes a MMAL_PARAMETER_RATIONAL_T
	ParameterClockLatency                                                 // Takes a MMAL_PARAMETER_CLOCK_LATENCY_T
)

type ParameterClockUpdateThresholdType struct {
	c C.MMAL_PARAMETER_CLOCK_UPDATE_THRESHOLD_T
}

func (p ParameterClockUpdateThresholdType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterClockUpdateThresholdType) Value() ClockUpdateThreshold {
	return ClockUpdateThreshold{p.c.value}
}

type ParameterClockDiscontThresholdType struct {
	c C.MMAL_PARAMETER_CLOCK_DISCONT_THRESHOLD_T
}

func (p ParameterClockDiscontThresholdType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterClockDiscontThresholdType) Value() ClockDiscontThreshold {
	return ClockDiscontThreshold{p.c.value}
}

type ParameterClockRequestThresholdType struct {
	c C.MMAL_PARAMETER_CLOCK_REQUEST_THRESHOLD_T
}

func (p ParameterClockRequestThresholdType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterClockRequestThresholdType) Value() ClockRequestThreshold {
	return ClockRequestThreshold{p.c.value}
}

type ParameterClockLatencyType struct {
	c C.MMAL_PARAMETER_CLOCK_LATENCY_T
}

func (p ParameterClockLatencyType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterClockLatencyType) Value() ClockLatency {
	return ClockLatency{p.c.value}
}
