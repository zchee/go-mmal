// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bcmhost

/*
#cgo pkg-config: bcm_host
#include <bcm_host.h>
*/
import "C"

func Init() {
	C.bcm_host_init()
}

func Deinit() {
	C.bcm_host_deinit()
}

func GraphicsGetDisplaySize(displayNumber uint16, width uint32, height uint32) int32 {
	w := C.uint32_t(width)
	h := C.uint32_t(height)
	return int32(C.graphics_get_display_size(C.uint16_t(displayNumber), &w, &h))
}

func GetPeripheralAddress() uint {
	return uint(C.bcm_host_get_peripheral_address()) 
}

func GetPeripheralSize() uint {
	return uint(C.bcm_host_get_peripheral_size()) 
}

func GetSdramAddress() uint {
	return uint(C.bcm_host_get_sdram_address()) 
}
