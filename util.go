// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/util/mmal_util.h>
*/
import "C"
import (
	"unsafe"
)

// TODO(zchee): implements
// #define MMAL_OFFSET(TYPE, FIELD) ((size_t)((uint8_t *)&((TYPE*)0)->FIELD - (uint8_t *)0))

func StatusToString(status Status) string {
	return C.GoString(C.mmal_status_to_string(C.MMAL_STATUS_T(status)))
}

func EncodingStrideToWidth(encoding uint32, stride uint32) uint32 {
	return uint32(C.mmal_encoding_stride_to_width(C.uint32_t(encoding), C.uint32_t(stride)))
}

func EncodingWidthToStride(encoding uint32, width uint32) uint32 {
	return uint32(C.mmal_encoding_width_to_stride(C.uint32_t(encoding), C.uint32_t(width)))
}

func EncodingGetSliceVariant(encoding uint32) uint32 {
	return uint32(C.mmal_encoding_get_slice_variant(C.uint32_t(encoding)))
}

func PortTypeToString(typ PortType) string {
	return C.GoString(C.mmal_port_type_to_string(C.MMAL_PORT_TYPE_T(typ)))
}

func PortParameterAllocGet(port *Port, id, size uint32) (ParameterHeader, error) {
	var s C.MMAL_STATUS_T
	phdr := C.mmal_port_parameter_alloc_get(port.c, C.uint32_t(id), C.uint32_t(size), &s)
	if Status(s) != Success {
		return ParameterHeader{}, Status(s)
	}
	return ParameterHeader{*phdr}, nil
}

func PortParameterFree(param *ParameterHeader) {
	C.mmal_port_parameter_free(&param.c)
}

func BufferHeaderCopyHeader(dest, src *BufferHeader) {
	C.mmal_buffer_header_copy_header(dest.c, src.c)
}

func PortPoolCreate(port *Port, headers uint, payloadSize uint32) PoolType {
	return PoolType{C.mmal_port_pool_create(port.c, C.uint(headers), C.uint32_t(payloadSize))}
}

func PortPoolDestroy(port *Port, pool *PoolType) {
	C.mmal_port_pool_destroy(port.c, pool.c)
}

func LogDumpPort(port *Port) {
	C.mmal_log_dump_port(port.c)
}

func LogDumpFormat(format *EsFormat) {
	C.mmal_log_dump_format(format.c)
}

func UtilGetPort(comp *ComponentType, typ PortType, index uint) Port {
	port := C.mmal_util_get_port(comp.c, C.MMAL_PORT_TYPE_T(typ), C.uint(index))
	return Port{port}
}

func FourccToString(buf string, len int, fourcc uint32) string {
	b := C.CString(buf)
	defer C.free(unsafe.Pointer(b))
	return C.GoString(C.mmal_4cc_to_string(b, C.size_t(len), C.uint32_t(fourcc)))
}

func UtilRgbOrderFixed(port *Port) int {
	return int(C.mmal_util_rgb_order_fixed(port.c))
}
