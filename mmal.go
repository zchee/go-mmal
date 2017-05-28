// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>

extern void set_es_specific_format(MMAL_ES_SPECIFIC_FORMAT_T *esf) {
	esf->video.width = 1024;
	esf->video.height = 768;
	esf->video.crop.x = 0;
	esf->video.crop.y = 0;
	esf->video.crop.width = 1024;
	esf->video.crop.height = 768;
	esf->video.frame_rate.num = 0;
	esf->video.frame_rate.den = 1;
}
*/
import "C"
import (
	"fmt"
)

func SetPortFormat(port *Port) {
	port.c.format.encoding = C.MMAL_ENCODING_OPAQUE

	var esf C.MMAL_ES_SPECIFIC_FORMAT_T
	// TODO(zchee): could not determine kind of name for C.VCOS_ALIGN_UP: C.VCOS_ALIGN_UP(1024, 32)
	// esf.video.width = 1024
	// TODO(zchee): could not determine kind of name for C.VCOS_ALIGN_UP: C.VCOS_ALIGN_UP(768, 16)
	// esf.video.height = 768
	// esf.video.crop.x = 0
	// esf.video.crop.y = 0
	// esf.video.crop.width = 1024
	// esf.video.crop.height = 768
	// esf.video.frame_rate.num = 0
	// esf.video.frame_rate.den = 1
	C.set_es_specific_format(&esf)
	port.c.format.es = &esf

	if err := PortFormatCommit(port); err != nil {
		panic(fmt.Sprintf("could not set port format: %v", err))
	}
}

func VcosSleep(wait uint32) {
	C.vcos_sleep(C.uint32_t(wait))
}
