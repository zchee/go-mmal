// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_buffer.h>
*/
import "C"
import "unsafe"

type BufferHeaderVideoSpecific struct {
	c C.MMAL_BUFFER_HEADER_VIDEO_SPECIFIC_T
}

func (b BufferHeaderVideoSpecific) Planes() uint32 {
	return uint32(b.c.planes)
}

func (b BufferHeaderVideoSpecific) Offset() []uint32 {
	sl := make([]uint32, len(b.c.offset))
	for _, i := range b.c.offset {
		sl = append(sl, uint32(i))
	}
	return sl
}

func (b BufferHeaderVideoSpecific) Pitch() []uint32 {
	sl := make([]uint32, len(b.c.pitch))
	for _, i := range b.c.pitch {
		sl = append(sl, uint32(i))
	}
	return sl
}

func (b BufferHeaderVideoSpecific) Flags() uint32 {
	return uint32(b.c.flags)
}

type BufferHeaderTypeSpecific C.MMAL_BUFFER_HEADER_TYPE_SPECIFIC_T

type BufferHeader struct {
	c *C.MMAL_BUFFER_HEADER_T
}

func (b BufferHeader) Next() BufferHeader {
	return BufferHeader{b.c.next}
}

// TODO(zchee): implements
func (b BufferHeader) Priv() {}

func (b BufferHeader) Cmd() uint32 {
	return uint32(b.c.cmd)
}

func (b BufferHeader) Data() uint8 {
	return uint8(*b.c.data)
}

func (b BufferHeader) AllocSize() uint32 {
	return uint32(b.c.alloc_size)
}

func (b BufferHeader) Length() uint32 {
	return uint32(b.c.length)
}

func (b BufferHeader) Offset() uint32 {
	return uint32(b.c.offset)
}

func (b BufferHeader) Flags() uint32 {
	return uint32(b.c.flags)
}

func (b BufferHeader) Pts() int64 {
	return int64(b.c.pts)
}

func (b BufferHeader) Dts() int64 {
	return int64(b.c.dts)
}

func (b BufferHeader) Type() BufferHeaderTypeSpecific {
	return BufferHeaderTypeSpecific(*b.c._type)
}

func (b BufferHeader) UserData() unsafe.Pointer {
	return b.c.user_data
}

const (
	// Signals that the current payload is the end of the stream of data
	BufferHeaderFlagEos = 1 << 0
	// Signals that the start of the current payload starts a frame
	BufferHeaderFlagFrameStart = 1 << 1
	// Signals that the end of the current payload ends a frame
	BufferHeaderFlagFrameEnd = 1 << 2
	// Signals that the current payload contains only complete frames = 1 or more
	BufferHeaderFlagFrame = BufferHeaderFlagFrameStart | BufferHeaderFlagFrameEnd
	// Signals that the current payload is a keyframe = i.e. self decodable
	BufferHeaderFlagKeyframe = 1 << 3
	// Signals a discontinuity in the stream of data = e.g. after a seek.
	// Can be used for instance by a decoder to reset its state
	BufferHeaderFlagDiscontinuity = 1 << 4
	// Signals a buffer containing some kind of config data for the component
	// = e.g. codec config data
	BufferHeaderFlagConfig = 1 << 5
	// Signals an encrypted payload
	BufferHeaderFlagEncrypted = 1 << 6
	// Signals a buffer containing side information
	BufferHeaderFlagCodecsideinfo = 1 << 7
	// Signals a buffer which is the snapshot/postview image from a stills capture
	BUFFER_HEADER_FlagsSnapshot = 1 << 8
	// Signals a buffer which contains data known to be corrupted
	BufferHeaderFlagCorrupted = 1 << 9
	// Signals that a buffer failed to be transmitted
	BufferHeaderFlagTransmission_failed = 1 << 10
	// Signals the output buffer won't be used, just update reference frames
	BufferHeaderFlagDecodeonly = 1 << 11
	// User flags - can be passed in and will get returned
	BufferHeaderFlagUser0 = 1 << 28
	BufferHeaderFlagUser1 = 1 << 29
	BufferHeaderFlagUser2 = 1 << 30
	BufferHeaderFlagUser3 = 1 << 31
)

const (
	BufferHeaderFlagFormatSpecificStart = 1 << 16
	// Signals an interlaced video frame.
	BufferHeaderVideoFlagInterlaced = BufferHeaderFlagFormatSpecificStart << 0
	// Signals that the top field of the current interlaced frame should be displayed first.
	BufferHeaderVideoFlagTopFieldFirst = BufferHeaderFlagFormatSpecificStart << 1
	// Signals that the buffer should be displayed on external display if attached.
	BufferHeaderVideoFlagDisplayExternal = BufferHeaderFlagFormatSpecificStart << 3
	// Signals that contents of the buffer requires copy protection.
	BufferHeaderVideoFlagProtected = BufferHeaderFlagFormatSpecificStart << 4
)

func BufferHeaderAcquire(header BufferHeader) {
	C.mmal_buffer_header_acquire(header.c)
}

func BufferHeaderReset(header BufferHeader) {
	C.mmal_buffer_header_reset(header.c)
}

func BufferHeaderRelease(header BufferHeader) {
	C.mmal_buffer_header_release(header.c)
}

func BufferHeaderReleseContinue(header BufferHeader) {
	C.mmal_buffer_header_release_continue(header.c)
}

// TODO(zchee): implements
// func BufferHeaderPreReleaseCBSet(header BufferHeader, cb ) {}

func BufferHeaderReplicate(dest, src BufferHeader) Status {
	return Status(C.mmal_buffer_header_replicate(dest.c, src.c))
}

func BufferHeaderMemLock(header BufferHeader) Status {
	return Status(C.mmal_buffer_header_mem_lock(header.c))
}

func BufferHeaderMemUnlock(header BufferHeader) {
	C.mmal_buffer_header_mem_unlock(header.c)
}
