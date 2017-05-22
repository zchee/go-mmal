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

func ComponentCreate(name string, component ComponentType) Status {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(&cName))
	return Status(C.mmal_component_create(cName, &component.c))
}

func ComponentAcquire(component ComponentType) {
	C.mmal_component_acquire(component.c)
}

func ComponentRelease(component ComponentType) Status {
	return Status(C.mmal_component_release(component.c))
}

func ComponentDestroy(component ComponentType) Status {
	return Status(C.mmal_component_destroy(component.c))
}

func ComponentEnable(component ComponentType) Status {
	return Status(C.mmal_component_enable(component.c))
}

func ComponentDisable(component ComponentType) Status {
	return Status(C.mmal_component_disable(component.c))
}
