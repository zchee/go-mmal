// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_parameters_audio.h>
*/
import "C"

const (
	ParameterAudioDestination   ParameterGroupAudio = (3 << 16) + iota // Takes a MMAL_PARAMETER_STRING_T
	ParameterAudioLatencyTarget                                        // Takes a MMAL_PARAMETER_AUDIO_LATENCY_TARGET_T
	ParameterAudioSource
	ParameterAudioPassthrough // Takes a MMAL_PARAMETER_BOOLEAN_T
)

type ParameterAudioLatencyTargetType struct {
	c C.MMAL_PARAMETER_AUDIO_LATENCY_TARGET_T
}

func (p ParameterAudioLatencyTargetType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterAudioLatencyTargetType) Enable() bool {
	return BoolT(p.c.enable) != BoolT(0)
}

func (p ParameterAudioLatencyTargetType) Filter() uint32 {
	return uint32(p.c.filter)
}

func (p ParameterAudioLatencyTargetType) Target() uint32 {
	return uint32(p.c.target)
}

func (p ParameterAudioLatencyTargetType) Shift() uint32 {
	return uint32(p.c.shift)
}

func (p ParameterAudioLatencyTargetType) SpeedFactor() uint32 {
	return uint32(p.c.speed_factor)
}

func (p ParameterAudioLatencyTargetType) InterFactor() uint32 {
	return uint32(p.c.inter_factor)
}

func (p ParameterAudioLatencyTargetType) AdjCap() uint32 {
	return uint32(p.c.adj_cap)
}
