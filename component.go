// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal_component.h>
*/
import "C"
import (
	"unsafe"
)

type ComponentType struct {
	c *C.MMAL_COMPONENT_T
}

type ComponentPrivateType struct {
	c *C.MMAL_COMPONENT_PRIVATE_T
}

func (c ComponentType) Priv() ComponentPrivateType {
	return ComponentPrivateType{c.c.priv}
}

type ComponentUserdataType struct {
	c *C.struct_MMAL_COMPONENT_USERDATA_T
}

func (c ComponentType) Userdata() ComponentUserdataType {
	return ComponentUserdataType{c.c.userdata}
}

func (c ComponentType) Name() string {
	return C.GoString(c.c.name)
}

func (c ComponentType) IsEnabled() uint32 {
	return uint32(c.c.is_enabled)
}

func (c ComponentType) Control() Port {
	return Port{c.c.control}
}

func (c ComponentType) InputNum() uint32 {
	return uint32(c.c.input_num)
}

func (c ComponentType) Input() Port {
	return Port{*c.c.input}
}

func (c ComponentType) OutputNum() uint32 {
	return uint32(c.c.output_num)
}

func (c ComponentType) Output() Port {
	return Port{*c.c.output}
}

func (c ComponentType) ClockNum() uint32 {
	return uint32(c.c.clock_num)
}

func (c ComponentType) Clock() Port {
	return Port{*c.c.clock}
}

func (c ComponentType) PortNum() uint32 {
	return uint32(c.c.port_num)
}

func (c ComponentType) Port() Port {
	return Port{*c.c.port}
}

func (c ComponentType) Id() uint32 {
	return uint32(c.c.id)
}

func ComponentCreate(name string, component *ComponentType) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	if err := Status(C.mmal_component_create(cName, &component.c)); err != Success {
		return err
	}
	return nil
}

func ComponentAcquire(component *ComponentType) {
	C.mmal_component_acquire(component.c)
}

func ComponentRelease(component *ComponentType) error {
	if err := Status(C.mmal_component_release(component.c)); err != Success {
		return err
	}
	return nil
}

func ComponentDestroy(component *ComponentType) error {
	if err := Status(C.mmal_component_destroy(component.c)); err != Success {
		return err
	}
	return nil
}

func ComponentEnable(component *ComponentType) error {
	if err := Status(C.mmal_component_enable(component.c)); err != Success {
		return err
	}
	return nil
}

func ComponentDisable(component *ComponentType) error {
	if err := Status(C.mmal_component_disable(component.c)); err != Success {
		return err
	}
	return nil
}
