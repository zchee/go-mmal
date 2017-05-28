// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_events.h>
*/
import "C"

var (
	// Error event. Data contains a MMAL_STATUS_T
	EventError = FourCC('E', 'R', 'R', 'O')
	// End-of-stream event. Data contains a MMAL_EVENT_END_OF_STREAM_T
	EventEOS = FourCC('E', 'E', 'O', 'S')
	// Format changed event. Data contains a MMAL_EVENT_FORMAT_CHANGED_T
	EventFormatChanged = FourCC('E', 'F', 'C', 'H')
	// Parameter changed event. Data contains the new parameter value, see MMAL_EVENT_PARAMETER_CHANGED_T
	EventParameterChanged = FourCC('E', 'P', 'C', 'H')
)

type EventEndOfStreamType struct {
	c C.MMAL_EVENT_END_OF_STREAM_T
}

func (e EventEndOfStreamType) PortType() PortType {
	return PortType(e.c.port_type)
}

func (e EventEndOfStreamType) PortIndex() uint32 {
	return uint32(e.c.port_index)
}

type EventFormatChangedType struct {
	c *C.MMAL_EVENT_FORMAT_CHANGED_T
}

func (e EventFormatChangedType) BufferSizeMin() uint32 {
	return uint32(e.c.buffer_size_min)
}

func (e EventFormatChangedType) BufferNumMin() uint32 {
	return uint32(e.c.buffer_num_min)
}

func (e EventFormatChangedType) BufferSizeRecommended() uint32 {
	return uint32(e.c.buffer_size_recommended)
}

func (e EventFormatChangedType) BufferNumRecommended() uint32 {
	return uint32(e.c.buffer_num_recommended)
}

func (e EventFormatChangedType) Format() EsFormat {
	return EsFormat{e.c.format}
}

type EventParameterChangedType struct {
	c C.MMAL_EVENT_PARAMETER_CHANGED_T
}

func (e EventParameterChangedType) Hdr() ParameterHeader {
	return ParameterHeader{e.c.hdr}
}

func EventFormatChangedGet(buffer *BufferHeader) EventFormatChangedType {
	return EventFormatChangedType{C.mmal_event_format_changed_get(buffer.c)}
}
