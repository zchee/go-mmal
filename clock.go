// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_clock.h>

extern MMAL_CLOCK_EVENT_T clock_event_init(uint32_t id) {
	MMAL_CLOCK_EVENT_T event;
	event.id = id;
	event.magic = MMAL_CLOCK_EVENT_MAGIC;
	event.buffer = NULL;
	event.padding0 = 0;
	// event.data = union{0};
	event.padding1 = 0;
	return event;
}
*/
import "C"
import (
	"bytes"
	"encoding/binary"
)

var (
	// Clock event magic
	ClockEventMagic = FourCC('C', 'K', 'L', 'M')

	// Clock reference update
	ClockEventReference = FourCC('C', 'R', 'E', 'F')

	// Clock state update
	ClockEventActive = FourCC('C', 'A', 'C', 'T')

	// Clock scale update
	ClockEventScale = FourCC('C', 'S', 'C', 'A')

	// Clock media-time update
	ClockEventTime = FourCC('C', 'T', 'I', 'M')

	// Clock update threshold
	ClockEventUpdateThreshold = FourCC('C', 'U', 'T', 'H')

	// Clock discontinuity threshold
	ClockEventDiscontThreshold = FourCC('C', 'D', 'T', 'H')

	// Clock request threshold
	ClockEventRequestThreshold = FourCC('C', 'R', 'T', 'H')

	// Buffer statistics
	ClockEventInputBufferInfo  = FourCC('C', 'I', 'B', 'I')
	ClockEventOutputBufferInfo = FourCC('C', 'O', 'B', 'I')

	// Clock latency setting
	ClockEventLatency = FourCC('C', 'L', 'A', 'T')

	// Clock event not valid
	ClockEventInvalid = 0
)

type ClockUpdateThreshold struct {
	c C.MMAL_CLOCK_UPDATE_THRESHOLD_T
}

func (c ClockUpdateThreshold) ThreshouldLower() int64 {
	return int64(c.c.threshold_lower)
}

func (c ClockUpdateThreshold) ThreshouldUpper() int64 {
	return int64(c.c.threshold_upper)
}

type ClockDiscontThreshold struct {
	c C.MMAL_CLOCK_DISCONT_THRESHOLD_T
}

func (c ClockDiscontThreshold) Threshold() int64 {
	return int64(c.c.threshold)
}

func (c ClockDiscontThreshold) Duration() int64 {
	return int64(c.c.duration)
}

type ClockRequestThreshold struct {
	c C.MMAL_CLOCK_REQUEST_THRESHOLD_T
}

func (c ClockRequestThreshold) Threshold() int64 {
	return int64(c.c.threshold)
}

func (c ClockRequestThreshold) ThresholdEnable() bool {
	return BoolT(c.c.threshold_enable) != BoolT(0)
}

type ClockBufferInfo struct {
	c C.MMAL_CLOCK_BUFFER_INFO_T
}

func (c ClockBufferInfo) TimeStamp() int64 {
	return int64(c.c.time_stamp)
}

func (c ClockBufferInfo) ArrivalTime() int64 {
	return int64(c.c.arrival_time)
}

type ClockLatency struct {
	c C.MMAL_CLOCK_LATENCY_T
}

func (c ClockLatency) Target() int64 {
	return int64(c.c.target)
}

func (c ClockLatency) AttackPeriod() int64 {
	return int64(c.c.attack_period)
}

func (c ClockLatency) AttackRate() int64 {
	return int64(c.c.attack_rate)
}

type ClockEvent struct {
	c C.MMAL_CLOCK_EVENT_T
}

func (c ClockEvent) ID() uint32 {
	return uint32(c.c.id)
}

func (c ClockEvent) Magic() uint32 {
	return uint32(c.c.magic)
}

func (c ClockEvent) Buffer() BufferHeader {
	return BufferHeader{c.c.buffer}
}

func (c ClockEvent) Padding0() uint32 {
	return uint32(c.c.padding0)
}

func (c ClockEvent) Data() *ClockEventData {
	buf := bytes.NewBuffer(c.c.data[:])
	var ptr *ClockEventData
	err := binary.Read(buf, binary.LittleEndian, &ptr)
	if err != nil {
		return nil
	}
	return ptr
}

func (c ClockEvent) Padding1() uint64 {
	return uint64(c.c.padding1)
}

func ClockEventInit(id uint32) ClockEvent {
	return ClockEvent{C.clock_event_init(C.uint32_t(id))}
}

type ClockEventData struct {
	// used either for clock reference or clock stateg
	enable C.MMAL_BOOL_T

	// new clock scaleg
	scale C.MMAL_RATIONAL_T

	// new media-timeg
	mediaTime C.int64_t

	// media-time update thresholdg
	updateThreshold C.MMAL_CLOCK_UPDATE_THRESHOLD_T

	// media-time discontinuity thresholdg
	discontThreshold C.MMAL_CLOCK_DISCONT_THRESHOLD_T

	// client callback request thresholdg
	requestThreshold C.MMAL_CLOCK_REQUEST_THRESHOLD_T

	// input/output buffer informationg
	buffer C.MMAL_CLOCK_BUFFER_INFO_T

	// clock latency settingg
	latency C.MMAL_CLOCK_LATENCY_T
}

func (c ClockEventData) Enable() bool {
	return BoolT(c.enable) != BoolT(0)
}

func (c ClockEventData) Scale() Rational {
	return Rational{c.scale}
}

func (c ClockEventData) MediaTime() int64 {
	return int64(c.mediaTime)
}

func (c ClockEventData) UpdateThreshold() ClockUpdateThreshold {
	return ClockUpdateThreshold{c.updateThreshold}
}

func (c ClockEventData) DiscontThreshold() ClockDiscontThreshold {
	return ClockDiscontThreshold{c.discontThreshold}
}

func (c ClockEventData) RequestThreshold() ClockRequestThreshold {
	return ClockRequestThreshold{c.requestThreshold}
}

func (c ClockEventData) Buffer() ClockBufferInfo {
	return ClockBufferInfo{c.buffer}
}

func (c ClockEventData) Latency() ClockLatency {
	return ClockLatency{c.latency}
}
