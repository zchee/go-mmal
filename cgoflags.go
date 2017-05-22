// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
// for local environment
// #cgo CFLAGS: -I/opt/vc/include -I/opt/vc/include/interface/vcos/pthreads
// #cgo LDFLAGS: -L/opt/vc/lib -lmmal -lmmal_core -lmmal_util -lmmal_vc_client -lbcm_host -lvcsm
// for raspberry pi native
#cgo pkg-config: mmal
*/
import "C"
