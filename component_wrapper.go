// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/util/mmal_component_wrapper.h>
*/
import "C"
import (
	"unsafe"
)

type WrapperType struct {
	c *C.MMAL_WRAPPER_T
}

type WrapperCallbackType struct {
	c *C.MMAL_WRAPPER_CALLBACK_T
}

func (w WrapperType) UserData() unsafe.Pointer {
	return w.c.user_data
}

func (w WrapperType) Callback() WrapperCallbackType {
	return WrapperCallbackType{&w.c.callback}
}

func (w WrapperType) Component() ComponentType {
	return ComponentType{w.c.component}
}

func (w WrapperType) Status() Status {
	return Status(w.c.status)
}

func (w WrapperType) Control() Port {
	return Port{w.c.control}
}

func (w WrapperType) InputNum() uint32 {
	return uint32(w.c.input_num)
}

func (w WrapperType) Input() Port {
	return Port{*w.c.input}
}

func (w WrapperType) InputPool() PoolType {
	return PoolType{*w.c.input_pool}
}

func (w WrapperType) OutputNum() uint32 {
	return uint32(w.c.output_num)
}

func (w WrapperType) Output() Port {
	return Port{*w.c.output}
}

func (w WrapperType) OutputPool() PoolType {
	return PoolType{*w.c.output_pool}
}

func (w WrapperType) OutputQueue() QueueType {
	return QueueType{*w.c.output_queue}
}

func (w WrapperType) TimeSetup() int64 {
	return int64(w.c.time_setup)
}

func (w WrapperType) TimeEnable() int64 {
	return int64(w.c.time_enable)
}

func (w WrapperType) TimeDisable() int64 {
	return int64(w.c.time_disable)
}

func WrapperCreate(wrapper *WrapperType, name string) Status {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return Status(C.mmal_wrapper_create((**C.MMAL_WRAPPER_T)(&wrapper.c), n))
}

const (
	// The operation should be blocking
	WrapperFlagWait = 1 << iota
	// The pool for the port should allocate memory for the payloads
	WrapperFlagPayloadAllocate
	// The port will use shared memory payloads
	WrapperFlagPayloadUseSharedMemory
)

func WrapperPortEnable(port *Port, flags uint32) Status {
	return Status(C.mmal_wrapper_port_enable((*C.MMAL_PORT_T)(unsafe.Pointer(&port.c)), C.uint32_t(flags)))
}

func WrapperPortDisable(port *Port) Status {
	return Status(C.mmal_wrapper_port_disable((*C.MMAL_PORT_T)(unsafe.Pointer(&port.c))))
}

func WrapperBufferGetEmpty(port *Port, buffer *BufferHeader, flags uint32) Status {
	return Status(C.mmal_wrapper_buffer_get_empty((*C.MMAL_PORT_T)(unsafe.Pointer(&port.c)), (**C.MMAL_BUFFER_HEADER_T)(unsafe.Pointer(&buffer.c)), C.uint32_t(flags)))
}

func WrapperBufferGetFull(port *Port, buffer *BufferHeader, flags uint32) Status {
	return Status(C.mmal_wrapper_buffer_get_full((*C.MMAL_PORT_T)(unsafe.Pointer(&port.c)), (**C.MMAL_BUFFER_HEADER_T)(unsafe.Pointer(&buffer.c)), C.uint32_t(flags)))
}

// TODO(zchee): mmal_wrapper_cancel does not appear to be in libmmal.so
// func WrapperCancel(wrapper WrapperType) Status {
// 	return Status(C.mmal_wrapper_cancel(&wrapper.c))
// }

func WrapperDestroy(wrapper *WrapperType) Status {
	return Status(C.mmal_wrapper_destroy(wrapper.c))
}
