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
// func (g GraphType) PfParameterSet() Status {}
// func (g GraphType) PfParameterGet() Status {}
// func (g GraphType) PfFormatCommit() Status {}
// func (g GraphType) PfSendBuffer() Status {}
// func (g GraphType) PfReturnBuffer() Status {}
// func (g GraphType) PfPayloadAlloc() Status {}
// func (g GraphType) PfPayloadFree() Status {}
// func (g GraphType) PfFlush() Status {}
// func (g GraphType) PfEnable() Status {}
// func (g GraphType) PfDisable() Status {}
// func (g GraphType) PfControlCallback() Status {}
// func (g GraphType) PfGraphEnable() Status {}
// func (g GraphType) PfConnectionBuffer() Status {}

func GraphCreate(graph *GraphType, userdataSize uint) Status {
	return Status(C.mmal_graph_create(&graph.c, C.uint(userdataSize)))
}

func GraphAddComponent(graph *GraphType, component *ComponentType) Status {
	return Status(C.mmal_graph_add_component(graph.c, component.c))
}

func GraphComponentTopology(graph *GraphType, component *ComponentType, topology GraphTopology, input int8, inputNum uint, output int8, outputNum uint) Status {
	i := C.int8_t(input)
	o := C.int8_t(output)
	return Status(C.mmal_graph_component_topology(graph.c, component.c, C.MMAL_GRAPH_TOPOLOGY_T(topology), &i, C.uint(inputNum), &o, C.uint(outputNum)))
}

func GraphAddPort(graph *GraphType, port *Port) Status {
	return Status(C.mmal_graph_add_port(graph.c, port.c))
}

func GraphAddConnection(graph *GraphType, connection *ConnectionType) Status {
	return Status(C.mmal_graph_add_connection(graph.c, connection.c))
}

func GraphNewComponent(graph *GraphType, name string, component *ComponentType) Status {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return Status(C.mmal_graph_new_component(graph.c, n, &component.c))
}

func GraphNewConnection(graph *GraphType, out *Port, in *Port, flags uint32, connection *ConnectionType) Status {
	return Status(C.mmal_graph_new_connection(graph.c, out.c, in.c, C.uint32_t(flags), &connection.c))
}

type GraphEventCb C.MMAL_GRAPH_EVENT_CB

func GraphEnable(graph *GraphType, cb GraphEventCb, cbData unsafe.Pointer) Status {
	return Status(C.mmal_graph_enable(graph.c, cb, cbData))
}

func GraphDisable(graph *GraphType) Status {
	return Status(C.mmal_graph_disable(graph.c))
}

func GraphFindPort(graph *GraphType, name string, typ PortType, index uint) Port {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return Port{C.mmal_graph_find_port(graph.c, n, C.MMAL_PORT_TYPE_T(typ), C.uint(index))}
}

func GraphBuild(ctx *GraphType, name string, component *ComponentType) Status {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return Status(C.mmal_graph_build(ctx.c, n, &component.c))
}

func GraphComponentConstructor(name string, component *ComponentType) Status {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return Status(C.mmal_graph_component_constructor(n, component.c))
}

func GraphDestroy(ctx *GraphType) Status {
	return Status(C.mmal_graph_destroy(ctx.c))
}
