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

func PortParameterSetBoolean(port *Port, id uint32, value BoolT) error {
	if err := Status(C.mmal_port_parameter_set_boolean(port.c, C.uint32_t(id), C.MMAL_BOOL_T(value))); err != Success {
		return err
	}
	return nil
}

func PortParameterGetBoolean(port *Port, id uint32, value *BoolT) error {
	v := C.MMAL_BOOL_T(*value)
	if err := Status(C.mmal_port_parameter_get_boolean(port.c, C.uint32_t(id), &v)); err != Success {
		return err
	}
	return nil
}

func PortParameterSetUint64(port *Port, id uint32, value uint64) error {
	if err := Status(C.mmal_port_parameter_set_uint64(port.c, C.uint32_t(id), C.uint64_t(value))); err != Success {
		return err
	}
	return nil
}

func PortParameterGetUint64(port *Port, id uint32, value *uint64) error {
	v := C.uint64_t(*value)
	if err := Status(C.mmal_port_parameter_get_uint64(port.c, C.uint32_t(id), &v)); err != Success {
		return err
	}
	return nil
}

func PortParameterSetInt64(port *Port, id uint32, value int64) error {
	if err := Status(C.mmal_port_parameter_set_int64(port.c, C.uint32_t(id), C.int64_t(value))); err != Success {
		return err
	}
	return nil
}

func PortParameterGetInt64(port *Port, id uint32, value *int64) error {
	v := C.int64_t(*value)
	if err := Status(C.mmal_port_parameter_get_int64(port.c, C.uint32_t(id), &v)); err != Success {
		return err
	}
	return nil
}

func PortParameterSetUint32(port *Port, id uint32, value uint32) error {
	if err := Status(C.mmal_port_parameter_set_uint32(port.c, C.uint32_t(id), C.uint32_t(value))); err != Success {
		return err
	}
	return nil
}

func PortParameterGetUint32(port *Port, id uint32, value *uint32) error {
	v := C.uint32_t(*value)
	if err := Status(C.mmal_port_parameter_get_uint32(port.c, C.uint32_t(id), &v)); err != Success {
		return err
	}
	return nil
}

func PortParameterSetInt32(port *Port, id uint32, value int32) error {
	if err := Status(C.mmal_port_parameter_set_int32(port.c, C.uint32_t(id), C.int32_t(value))); err != Success {
		return err
	}
	return nil
}

func PortParameterGetInt32(port *Port, id uint32, value *int32) error {
	v := C.int32_t(*value)
	if err := Status(C.mmal_port_parameter_get_int32(port.c, C.uint32_t(id), &v)); err != Success {
		return err
	}
	return nil
}

func PortParameterSetRational(port *Port, id uint32, value Rational) error {
	if err := Status(C.mmal_port_parameter_set_rational(port.c, C.uint32_t(id), C.MMAL_RATIONAL_T(value.c))); err != Success {
		return err
	}
	return nil
}

func PortParameterGetRational(port *Port, id uint32, value *Rational) error {
	if err := Status(C.mmal_port_parameter_get_rational(port.c, C.uint32_t(id), &value.c)); err != Success {
		return err
	}
	return nil
}

func PortParameterSetString(port *Port, id uint32, value string) error {
	v := C.CString(value)
	defer C.free(unsafe.Pointer(v))
	if err := Status(C.mmal_port_parameter_set_string(port.c, C.uint32_t(id), v)); err != Success {
		return err
	}
	return nil
}

func PortParameterSetBytes(port *Port, id uint32, data uint8, size uint) error {
	d := C.uint8_t(data)
	if err := Status(C.mmal_port_parameter_set_bytes(port.c, C.uint32_t(id), &d, C.uint(size))); err != Success {
		return err
	}
	return nil
}

func UtilPortSetUri(port *Port, uri string) error {
	u := C.CString(uri)
	defer C.free(unsafe.Pointer(u))
	if err := Status(C.mmal_util_port_set_uri(port.c, u)); err != Success {
		return err
	}
	return nil
}

func UtilSetDisplayRegion(port *Port, region *Displayregion) error {
	if err := Status(C.mmal_util_set_display_region(port.c, &region.c)); err != Success {
		return err
	}
	return nil
}

func UtilCameraUseStcTimestamp(port *Port, mode CameraSTCMode) error {
	if err := Status(C.mmal_util_camera_use_stc_timestamp(port.c, C.MMAL_CAMERA_STC_MODE_T(mode))); err != Success {
		return err
	}
	return nil
}

func UtilGetCorePortStats(port *Port, dir CoreStatsDir, reset BoolT, stats CoreStatisticsType) error {
	if err := Status(C.mmal_util_get_core_port_stats(port.c, C.MMAL_CORE_STATS_DIR(dir), C.MMAL_BOOL_T(reset), &stats.c)); err != Success {
		return err
	}
	return nil
}
