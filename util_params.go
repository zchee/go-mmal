// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/util/mmal_util_params.h>
*/
import "C"
import (
	"unsafe"
)

func PortParameterSetBoolean(port *Port, id uint32, value BoolT) Status {
	return Status(C.mmal_port_parameter_set_boolean(port.c, C.uint32_t(id), C.MMAL_BOOL_T(value)))
}

func PortParameterGetBoolean(port *Port, id uint32, value *BoolT) Status {
	v := C.MMAL_BOOL_T(*value)
	return Status(C.mmal_port_parameter_get_boolean(port.c, C.uint32_t(id), &v))
}

func PortParameterSetUint64(port *Port, id uint32, value uint64) Status {
	return Status(C.mmal_port_parameter_set_uint64(port.c, C.uint32_t(id), C.uint64_t(value)))
}

func PortParameterGetUint64(port *Port, id uint32, value *uint64) Status {
	v := C.uint64_t(*value)
	return Status(C.mmal_port_parameter_get_uint64(port.c, C.uint32_t(id), &v))
}

func PortParameterSetInt64(port *Port, id uint32, value int64) Status {
	return Status(C.mmal_port_parameter_set_int64(port.c, C.uint32_t(id), C.int64_t(value)))
}

func PortParameterGetInt64(port *Port, id uint32, value *int64) Status {
	v := C.int64_t(*value)
	return Status(C.mmal_port_parameter_get_int64(port.c, C.uint32_t(id), &v))
}

func PortParameterSetUint32(port *Port, id uint32, value uint32) Status {
	return Status(C.mmal_port_parameter_set_uint32(port.c, C.uint32_t(id), C.uint32_t(value)))
}

func PortParameterGetUint32(port *Port, id uint32, value *uint32) Status {
	v := C.uint32_t(*value)
	return Status(C.mmal_port_parameter_get_uint32(port.c, C.uint32_t(id), &v))
}

func PortParameterSetInt32(port *Port, id uint32, value int32) Status {
	return Status(C.mmal_port_parameter_set_int32(port.c, C.uint32_t(id), C.int32_t(value)))
}

func PortParameterGetInt32(port *Port, id uint32, value *int32) Status {
	v := C.int32_t(*value)
	return Status(C.mmal_port_parameter_get_int32(port.c, C.uint32_t(id), &v))
}

func PortParameterSetRational(port *Port, id uint32, value Rational) Status {
	return Status(C.mmal_port_parameter_set_rational(port.c, C.uint32_t(id), C.MMAL_RATIONAL_T(value.c)))
}

func PortParameterGetRational(port *Port, id uint32, value *Rational) Status {
	return Status(C.mmal_port_parameter_get_rational(port.c, C.uint32_t(id), &value.c))
}

func PortParameterSetString(port *Port, id uint32, value string) Status {
	v := C.CString(value)
	defer C.free(unsafe.Pointer(&v))
	return Status(C.mmal_port_parameter_set_string(port.c, C.uint32_t(id), v))
}

func PortParameterSetBytes(port *Port, id uint32, data uint8, size uint) Status {
	d := C.uint8_t(data)
	return Status(C.mmal_port_parameter_set_bytes(port.c, C.uint32_t(id), &d, C.uint(size)))
}

func UtilPortSetUri(port *Port, uri string) Status {
	u := C.CString(uri)
	defer C.free(unsafe.Pointer(&u))
	return Status(C.mmal_util_port_set_uri(port.c, u))
}

func UtilSetDisplayRegion(port *Port, region *Displayregion) Status {
	return Status(C.mmal_util_set_display_region(port.c, &region.c))
}

func UtilCameraUseStcTimestamp(port *Port, mode CameraSTCMode) Status {
	return Status(C.mmal_util_camera_use_stc_timestamp(port.c, C.MMAL_CAMERA_STC_MODE_T(mode)))
}

func UtilGetCorePortStats(port *Port, dir CoreStatsDir, reset BoolT, stats CoreStatisticsType) Status {
	return Status(C.mmal_util_get_core_port_stats(port.c, C.MMAL_CORE_STATS_DIR(dir), C.MMAL_BOOL_T(reset), &stats.c))
}
