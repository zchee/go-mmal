// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_buffer.h>
#include <interface/mmal/core/mmal_buffer_private.h>
*/
import "C"
import (
	"unsafe"
)

const DriverBufferSize = 32

type DriverBufferType struct {
	c *C.MMAL_DRIVER_BUFFER_T
}

type BufferHeaderPrivateType struct {
	c *C.MMAL_BUFFER_HEADER_PRIVATE_T
}

// TODO(zchee): implements. error: type *_Ctype_struct_MMAL_BUFFER_HEADER_PRIVATE_T has no field or method ***(e.g. pf_pre_release)
// func (b BufferHeaderPrivateType) PfPreRelease() BHPreReleaseCB {
// 	return BHPreReleaseCB(b.c.pf_pre_release)
// }
//
//
// func (b BufferHeaderPrivateType) PreReleaseUserdata() unsafe.Pointer {
// 	return b.c.pre_release_userdata
// }
//
// func (b BufferHeaderPrivateType) PfRelease() Void (*)(mmalBufferHeaderT) {
// 	return Void (*)(mmalBufferHeaderT){b.c.pf_release}
// }
//
// func (b BufferHeaderPrivateType) Owner() unsafe.Pointer {
// 	return b.c.owner
// }
//
// func (b BufferHeaderPrivateType) Refcount() int32 {
// 	return int32(b.c.refcount)
// }
//
// func (b BufferHeaderPrivateType) Reference() BufferHeader {
// 	return BufferHeader{b.c.reference}
// }
//
// func (b BufferHeaderPrivateType) PfPayloadFree() Void (*)(void, void *) {
// 	return Void (*)(void, void *){b.c.pf_payload_free}
// }
//
// func (b BufferHeaderPrivateType) Payload() unsafe.Pointer {
// 	return b.c.payload
// }
//
// func (b BufferHeaderPrivateType) PayloadContext() unsafe.Pointer {
// 	return b.c.payload_context
// }
//
// func (b BufferHeaderPrivateType) PayloadSize() Status {
// 	return Status{b.c.payload_size}
// }
//
// func (b BufferHeaderPrivateType) ComponentData() unsafe.Pointer {
// 	return b.c.component_data
// }
//
// func (b BufferHeaderPrivateType) PayloadHandle() unsafe.Pointer {
// 	return b.c.payload_handle
// }
//
// func (b BufferHeaderPrivateType) DriverArea() []uint {
// 	darea := make([]uint, DriverBufferSize)
// 	for i, da := range b.c.driver_area {
// 		darea[i] = da
// 	}
// 	return darea
// }

func BufferHeaderSize(header *BufferHeader) uint {
	return uint(C.mmal_buffer_header_size(header.c))
}

func BufferHeaderInitialise(mem unsafe.Pointer, length uint) BufferHeader {
	return BufferHeader{C.mmal_buffer_header_initialise(mem, C.uint(length))}
}

func BufferHeaderDriverData(header *BufferHeader) DriverBufferType {
	return DriverBufferType{C.mmal_buffer_header_driver_data(header.c)}
}

func BufferHeaderReference(header *BufferHeader) BufferHeader {
	return BufferHeader{C.mmal_buffer_header_reference(header.c)}
}
