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

// TODO(zchee): implements ComponentType methods. error: undefined several C side methods.

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
