// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/util/mmal_connection.h>
*/
import "C"
import "unsafe"

const (
	// The connection is tunnelled. Buffer headers do not transit via the client but directly from the output port to the input port. */
	ConnectionFlagTunnelling = 0x1
	// Force the pool of buffer headers used by the connection to be allocated on the input port. */
	ConnectionFlagAllocationOnInput = 0x2
	// Force the pool of buffer headers used by the connection to be allocated on the output port. */
	ConnectionFlagAllocationOnOutput = 0x4
	// Specify that the connection should not modify the buffer requirements. */
	ConnectionFlagKeepBufferRequirements = 0x8
	// The connection is flagged as direct. This doesn't change the behaviour of the connection itself but is used by the the graph utility to specify that the buffer should be sent to the input port from with the port callback. */
	ConnectionFlagDirect = 0x10
)

type ConnectionType struct {
	c C.MMAL_CONNECTION_T
}

func (c ConnectionType) UserData() unsafe.Pointer {
	return c.c.user_data
}

func (c ConnectionType) Callback() ConnectionCallbackType {
	return ConnectionCallbackType{c.c.callback}
}

type ConnectionCallbackType struct {
	c C.MMAL_CONNECTION_CALLBACK_T
}

func (c ConnectionType) IsEnabled() uint32 {
	return uint32(c.c.is_enabled)
}

func (c ConnectionType) Flags() uint32 {
	return uint32(c.c.flags)
}

func (c ConnectionType) In() Port {
	return Port{*c.c.in}
}

func (c ConnectionType) Out() Port {
	return Port{*c.c.out}
}

func (c ConnectionType) Pool() PoolType {
	return PoolType{c.c.pool}
}

func (c ConnectionType) Queue() QueueType {
	return QueueType{c.c.queue}
}

func (c ConnectionType) Name() string {
	return C.GoString(c.c.name)
}

func (c ConnectionType) TimeSetup() int64 {
	return int64(c.c.time_setup)
}

func (c ConnectionType) TimeEnable() int64 {
	return int64(c.c.time_enable)
}

func (c ConnectionType) TimeDisable() int64 {
	return int64(c.c.time_disable)
}

func ConnectionCreate(connection ConnectionType, out, in Port, flags uint32) Status {
	conn := &connection.c
	return Status(C.mmal_connection_create(&conn, &out.c, &in.c, C.uint32_t(flags)))
}

func ConnectionAcquire(connection ConnectionType) {
	C.mmal_connection_acquire(&connection.c)
}

func ConnectionRelease(connection ConnectionType) Status {
	return Status(C.mmal_connection_release(&connection.c))
}

func ConnectionDestroy(connection ConnectionType) Status {
	return Status(C.mmal_connection_destroy(&connection.c))
}

func ConnectionEnable(connection ConnectionType) Status {
	return Status(C.mmal_connection_enable(&connection.c))
}

func ConnectionDisable(connection ConnectionType) Status {
	return Status(C.mmal_connection_disable(&connection.c))
}

func ConnectionEventFormatChanged(connection ConnectionType, buffer BufferHeader) Status {
	return Status(C.mmal_connection_event_format_changed(&connection.c, buffer.c))
}
