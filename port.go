// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal_port.h>
*/
import "C"
import "unsafe"

type PortType C.MMAL_PORT_TYPE_T

const (
	PortTypeUnknown PortType = iota
	PortTypeControl
	PortTypeInput
	PortTypeOutput
	PortTypeClock
	PortTypeInvalid = 0xffffffff
)

const (
	PortCapabilityPassthrough               = 0x01
	PortCapabilityAllocation                = 0x02
	PortCapabilitySupportsEventFormatChange = 0x04
)

type Port struct {
	c *C.MMAL_PORT_T
}

func (p Port) Priv() PortPrivateType {
	return PortPrivateType{p.c.priv}
}

func (p Port) Name() string {
	return C.GoString(p.c.name)
}

func (p Port) Type() PortType {
	return PortType(p.c._type)
}

func (p Port) Index() uint16 {
	return uint16(p.c.index)
}

func (p Port) IndexAll() uint16 {
	return uint16(p.c.index_all)
}

func (p Port) IsEnabled() uint32 {
	return uint32(p.c.is_enabled)
}

func (p Port) Format() EsFormat {
	return EsFormat{p.c.format}
}

func (p Port) BufferNumMin() uint32 {
	return uint32(p.c.buffer_num_min)
}

func (p Port) BufferSizeMin() uint32 {
	return uint32(p.c.buffer_size_min)
}

func (p Port) BufferAlignmentMin() uint32 {
	return uint32(p.c.buffer_alignment_min)
}

func (p Port) BufferNumRecommended() uint32 {
	return uint32(p.c.buffer_num_recommended)
}

func (p Port) BufferSizeRecommended() uint32 {
	return uint32(p.c.buffer_size_recommended)
}

func (p Port) BufferNum() uint32 {
	return uint32(p.c.buffer_num)
}

func (p Port) BufferSize() uint32 {
	return uint32(p.c.buffer_size)
}

func (p Port) Component() ComponentType {
	return ComponentType{p.c.component}
}

type PortUserdataType struct {
	c *C.struct_MMAL_PORT_USERDATA_T
}

func (p Port) Userdata() PortUserdataType {
	return PortUserdataType{p.c.userdata}
}

func (p Port) Capabilities() uint32 {
	return uint32(p.c.capabilities)
}

func PortFormatCommit(port *Port) error {
	if err := Status(C.mmal_port_format_commit(port.c)); err != Success {
		return err
	}
	return nil
}

type PortBHCBType C.MMAL_PORT_BH_CB_T

func PortEnable(port Port, cb PortBHCBType) error {
	if err := Status(C.mmal_port_enable(port.c, cb)); err != Success {
		return err
	}
	return nil
}

func PortDisable(port *Port) error {
	if err := Status(C.mmal_port_disable(port.c)); err != Success {
		return err
	}
	return nil
}

func PortFlush(port *Port) error {
	if err := Status(C.mmal_port_flush(port.c)); err != Success {
		return err
	}
	return nil
}

func PortParameterSet(port *Port, param ParameterHeader) error {
	if err := Status(C.mmal_port_parameter_set(port.c, &param.c)); err != Success {
		return err
	}
	return nil
}

func PortParameterGet(port *Port, param ParameterHeader) error {
	if err := Status(C.mmal_port_parameter_get(port.c, &param.c)); err != Success {
		return err
	}
	return nil
}

func PortSendBuffer(port *Port, buffer BufferHeader) error {
	if err := Status(C.mmal_port_send_buffer(port.c, buffer.c)); err != Success {
		return err
	}
	return nil
}

func PortConnect(port, otherPort *Port) error {
	if err := Status(C.mmal_port_connect(port.c, otherPort.c)); err != Success {
		return err
	}
	return nil
}

func PortDisconnect(port *Port) error {
	if err := Status(C.mmal_port_disconnect(port.c)); err != Success {
		return err
	}
	return nil
}

// TODO(zchee): return type (*C.uint8)
func PortPayloadAlloc(port *Port, payloadSize uint32) unsafe.Pointer {
	res := C.mmal_port_payload_alloc(port.c, C.uint32_t(payloadSize))
	return unsafe.Pointer(res)
}

func PortPayloadFree(port *Port, payload *uint8) {
	p := C.uint8_t(*payload)
	C.mmal_port_payload_free(port.c, &p)
}

func PortEventGet(port *Port, buffer *BufferHeader, event uint32) error {
	if err := Status(C.mmal_port_event_get(port.c, &buffer.c, C.uint32_t(event))); err != Success {
		return err
	}
	return nil
}
