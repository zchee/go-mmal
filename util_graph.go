// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#cgo CFLAGS: -I/opt/vc/include/interface/mmal
#include <interface/mmal/mmal.h>
#include <interface/mmal/util/mmal_graph.h>
*/
import "C"
import (
	"unsafe"
)

type GraphTopology C.MMAL_GRAPH_TOPOLOGY_T

const (
	GraphTopologyAll      GraphTopology = C.MMAL_GRAPH_TOPOLOGY_ALL
	GraphTopologyStraight GraphTopology = C.MMAL_GRAPH_TOPOLOGY_STRAIGHT
	GraphTopologyCustom   GraphTopology = C.MMAL_GRAPH_TOPOLOGY_CUSTOM
	GraphTopologyMax      GraphTopology = C.MMAL_GRAPH_TOPOLOGY_MAX
)

type GraphType struct {
	c *C.MMAL_GRAPH_T
}

type GraphUserdataType struct {
	c *C.struct_MMAL_GRAPH_USERDATA_T
}

func (g GraphType) Userdata() GraphUserdataType {
	return GraphUserdataType{g.c.userdata}
}

// TODO(zchee): implements
// func (g GraphType) PfDestroy() GraphType {}
// func (g GraphType) PfParameterSet() error {}
// func (g GraphType) PfParameterGet() error {}
// func (g GraphType) PfFormatCommit() error {}
// func (g GraphType) PfSendBuffer() error {}
// func (g GraphType) PfReturnBuffer() error {}
// func (g GraphType) PfPayloadAlloc() error {}
// func (g GraphType) PfPayloadFree() error {}
// func (g GraphType) PfFlush() error {}
// func (g GraphType) PfEnable() error {}
// func (g GraphType) PfDisable() error {}
// func (g GraphType) PfControlCallback() error {}
// func (g GraphType) PfGraphEnable() error {}
// func (g GraphType) PfConnectionBuffer() error {}

func GraphCreate(graph *GraphType, userdataSize uint) error {
	if err := Status(C.mmal_graph_create(&graph.c, C.uint(userdataSize))); err != Success {
		return err
	}
	return nil
}

func GraphAddComponent(graph *GraphType, component *ComponentType) error {
	if err := Status(C.mmal_graph_add_component(graph.c, component.c)); err != Success {
		return err
	}
	return nil
}

func GraphComponentTopology(graph *GraphType, component *ComponentType, topology GraphTopology, input int8, inputNum uint, output int8, outputNum uint) error {
	i := C.int8_t(input)
	o := C.int8_t(output)
	if err := Status(C.mmal_graph_component_topology(graph.c, component.c, C.MMAL_GRAPH_TOPOLOGY_T(topology), &i, C.uint(inputNum), &o, C.uint(outputNum))); err != Success {
		return err
	}
	return nil
}

func GraphAddPort(graph *GraphType, port *Port) error {
	if err := Status(C.mmal_graph_add_port(graph.c, port.c)); err != Success {
		return err
	}
	return nil
}

func GraphAddConnection(graph *GraphType, connection *ConnectionType) error {
	if err := Status(C.mmal_graph_add_connection(graph.c, connection.c)); err != Success {
		return err
	}
	return nil
}

func GraphNewComponent(graph *GraphType, name string, component *ComponentType) error {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	if err := Status(C.mmal_graph_new_component(graph.c, n, &component.c)); err != Success {
		return err
	}
	return nil
}

func GraphNewConnection(graph *GraphType, out *Port, in *Port, flags uint32, connection *ConnectionType) error {
	if err := Status(C.mmal_graph_new_connection(graph.c, out.c, in.c, C.uint32_t(flags), &connection.c)); err != Success {
		return err
	}
	return nil
}

type GraphEventCb C.MMAL_GRAPH_EVENT_CB

func GraphEnable(graph *GraphType, cb GraphEventCb, cbData unsafe.Pointer) error {
	if err := Status(C.mmal_graph_enable(graph.c, cb, cbData)); err != Success {
		return err
	}
	return nil
}

func GraphDisable(graph *GraphType) error {
	if err := Status(C.mmal_graph_disable(graph.c)); err != Success {
		return err
	}
	return nil
}

func GraphFindPort(graph *GraphType, name string, typ PortType, index uint) Port {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return Port{C.mmal_graph_find_port(graph.c, n, C.MMAL_PORT_TYPE_T(typ), C.uint(index))}
}

func GraphBuild(ctx *GraphType, name string, component *ComponentType) error {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	if err := Status(C.mmal_graph_build(ctx.c, n, &component.c)); err != Success {
		return err
	}
	return nil
}

func GraphComponentConstructor(name string, component *ComponentType) error {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	if err := Status(C.mmal_graph_component_constructor(n, component.c)); err != Success {
		return err
	}
	return nil
}

func GraphDestroy(ctx *GraphType) error {
	if err := Status(C.mmal_graph_destroy(ctx.c)); err != Success {
		return err
	}
	return nil
}
