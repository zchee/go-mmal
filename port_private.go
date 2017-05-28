// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/core/mmal_port_private.h>
*/
import "C"
import (
	"unsafe"
)

type PortPrivateType struct {
	c *C.MMAL_PORT_PRIVATE_T
}

type PortPrivateCoreType struct {
	c *C.struct_MMAL_PORT_PRIVATE_CORE_T
}

type PortModuleType struct {
	c *C.struct_MMAL_PORT_MODULE_T
}

type PortClockType struct {
	c *C.struct_MMAL_PORT_CLOCK_T
}

// TODO(zchee): implements
// func (p PortPrivateType) Core() PortPrivateCoreType {
// 	return PortPrivateCoreType{p.c._core}
// }
//
// func (p PortPrivateType) Module() PortModuleType {
// 	return PortModuleType{p.c._module}
// }
//
// func (p PortPrivateType) Clock() PortClockType {
// 	return PortClockType{p.c._clock}
// }
//
// func (p PortPrivateType) PfSetFormat() StatusT (*)(mmalPortT) {
// 	return StatusT (*)(mmalPortT){p.c.pf_set_format}
// }
//
// func (p PortPrivateType) PfEnable() StatusT (*)(mmalPortT, mmalPortBhCbT) {
// 	return StatusT (*)(mmalPortT, mmalPortBhCbT){p.c.pf_enable}
// }
//
// func (p PortPrivateType) PfDisable() StatusT (*)(mmalPortT) {
// 	return StatusT (*)(mmalPortT){p.c.pf_disable}
// }
//
// func (p PortPrivateType) PfSend() StatusT (*)(mmalPortT, mmalBufferHeaderT *) {
// 	return StatusT (*)(mmalPortT, mmalBufferHeaderT *){p.c.pf_send}
// }
//
// func (p PortPrivateType) PfFlush() StatusT (*)(mmalPortT) {
// 	return StatusT (*)(mmalPortT){p.c.pf_flush}
// }
//
// func (p PortPrivateType) PfParameterSet() StatusT (*)(mmalPortT, const mmalParameterHeaderT *) {
// 	return StatusT (*)(mmalPortT, const mmalParameterHeaderT *){p.c.pf_parameter_set}
// }
//
// func (p PortPrivateType) PfParameterGet() StatusT (*)(mmalPortT, mmalParameterHeaderT *) {
// 	return StatusT (*)(mmalPortT, mmalParameterHeaderT *){p.c.pf_parameter_get}
// }
//
// func (p PortPrivateType) PfConnect() StatusT (*)(mmalPortT, mmalPortT *) {
// 	return StatusT (*)(mmalPortT, mmalPortT *){p.c.pf_connect}
// }
//
// func (p PortPrivateType) PfPayloadAlloc() Uint8T(*)(mmalPortT *, uint32T) {
// 	return Uint8T(*)(mmalPortT *, uint32T){p.c.pf_payload_alloc}
// }
//
// func (p PortPrivateType) PfPayloadFree() Void (*)(mmalPortT, uint8T *) {
// 	return Void (*)(mmalPortT, uint8T *){p.c.pf_payload_free}
// }

func PortBufferHeaderCallback(port *Port, buffer *BufferHeader) {
	C.mmal_port_buffer_header_callback(port.c, buffer.c)
}

func PortEventSend(port *Port, buffer *BufferHeader) {
	C.mmal_port_event_send(port.c, buffer.c)
}

func PortAlloc(component *ComponentType, typ PortType, extraSize uint) Port {
	return Port{C.mmal_port_alloc(component.c, C.MMAL_PORT_TYPE_T(typ), C.uint(extraSize))}
}

func PortFree(port *Port) {
	C.mmal_port_free(port.c)
}

func PortsAlloc(component *ComponentType, portsNum uint, typ PortType, extraSize uint) Port {
	port := C.mmal_ports_alloc(component.c, C.uint(portsNum), C.MMAL_PORT_TYPE_T(typ), C.uint(extraSize))
	return Port{*port}
}

func PortsFree(ports *Port, portsNum uint) {
	C.mmal_ports_free(&ports.c, C.uint(portsNum))
}

func PortAcquire(port *Port) {
	C.mmal_port_acquire(port.c)
}

func PortRelease(port *Port) error {
	if err := Status(C.mmal_port_release(port.c)); err != Success {
		return err
	}
	return nil
}

func PortPause(port *Port, pause BoolT) error {
	if err := Status(C.mmal_port_pause(port.c, C.MMAL_BOOL_T(pause))); err != Success {
		return err
	}
	return nil
}

func PortIsConnected(port *Port) bool {
	return BoolT(C.mmal_port_is_connected(port.c)) != BoolT(0)
}

type PortClockEventCB C.MMAL_PORT_CLOCK_EVENT_CB

func PortClockAlloc(component *ComponentType, extraSize uint, eventCb PortClockEventCB) Port {
	return Port{C.mmal_port_clock_alloc(component.c, C.uint(extraSize), eventCb)}
}

func PortClockFree(port *Port) {
	C.mmal_port_clock_free(port.c)
}

func PortsClockAlloc(component *ComponentType, portsNum, extraSize uint, eventCb PortClockEventCB) Port {
	ports := C.mmal_ports_clock_alloc(component.c, C.uint(portsNum), C.uint(extraSize), eventCb)
	return Port{*ports}
}

func PortsClockFree(ports *Port, portsNum uint) {
	C.mmal_ports_clock_free(&ports.c, C.uint(portsNum))
}

type PortClockRequestCB C.MMAL_PORT_CLOCK_REQUEST_CB

func PortClockRequestAdd(port *Port, mediaTime int64, cb PortClockRequestCB, cbData unsafe.Pointer) error {
	if err := Status(C.mmal_port_clock_request_add(port.c, C.int64_t(mediaTime), cb, cbData)); err != Success {
		return err
	}
	return nil
}

func PortClockRequestFlush(port *Port) error {
	if err := Status(C.mmal_port_clock_request_flush(port.c)); err != Success {
		return err
	}
	return nil
}

func PortClockReferenceSet(port *Port, reference BoolT) error {
	if err := Status(C.mmal_port_clock_reference_set(port.c, C.MMAL_BOOL_T(reference))); err != Success {
		return err
	}
	return nil
}

func PortClockReferenceGet(port *Port) bool {
	return BoolT(C.mmal_port_clock_reference_get(port.c)) != BoolT(0)
}

func PortClockActiveSet(port *Port, active BoolT) error {
	if err := Status(C.mmal_port_clock_active_set(port.c, C.MMAL_BOOL_T(active))); err != Success {
		return err
	}
	return nil
}

func PortClockActiveGet(port *Port) bool {
	return BoolT(C.mmal_port_clock_active_get(port.c)) != BoolT(0)
}

func PortClockScaleSet(port *Port, scale Rational) error {
	if err := Status(C.mmal_port_clock_scale_set(port.c, scale.c)); err != Success {
		return err
	}
	return nil
}

func PortClockScaleGet(port *Port) Rational {
	return Rational{C.mmal_port_clock_scale_get(port.c)}
}

func PortClockMediaTimeSet(port *Port, mediaTime int64) error {
	if err := Status(C.mmal_port_clock_media_time_set(port.c, C.int64_t(mediaTime))); err != Success {
		return err
	}
	return nil
}

func PortClockMediaTimeGet(port *Port) int64 {
	return int64(C.mmal_port_clock_media_time_get(port.c))
}

func PortClockUpdateThresholdSet(port *Port, threshold *ClockUpdateThreshold) error {
	if err := Status(C.mmal_port_clock_update_threshold_set(port.c, &threshold.c)); err != Success {
		return err
	}
	return nil
}

func PortClockUpdateThresholdGet(port *Port, threshold *ClockUpdateThreshold) error {
	if err := Status(C.mmal_port_clock_update_threshold_get(port.c, &threshold.c)); err != Success {
		return err
	}
	return nil
}

func PortClockDiscontThresholdSet(port *Port, threshold *ClockDiscontThreshold) error {
	if err := Status(C.mmal_port_clock_discont_threshold_set(port.c, &threshold.c)); err != Success {
		return err
	}
	return nil
}

func PortClockDiscontThresholdGet(port *Port, threshold *ClockDiscontThreshold) error {
	if err := Status(C.mmal_port_clock_discont_threshold_get(port.c, &threshold.c)); err != Success {
		return err
	}
	return nil
}

func PortClockRequestThresholdSet(port *Port, threshold *ClockRequestThreshold) error {
	if err := Status(C.mmal_port_clock_request_threshold_set(port.c, &threshold.c)); err != Success {
		return err
	}
	return nil
}

func PortClockRequestThresholdGet(port *Port, threshold *ClockRequestThreshold) error {
	if err := Status(C.mmal_port_clock_request_threshold_get(port.c, &threshold.c)); err != Success {
		return err
	}
	return nil
}

func PortClockInputBufferInfo(port *Port, info *ClockBufferInfo) {
	C.mmal_port_clock_input_buffer_info(port.c, &info.c)
}

func PortClockOutputBufferInfo(port *Port, info *ClockBufferInfo) {
	C.mmal_port_clock_output_buffer_info(port.c, &info.c)
}
