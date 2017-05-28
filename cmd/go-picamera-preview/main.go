// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Copyright (c) 2017 Sugizaki Yukimasa (ysugi@idein.jp)
// All rights reserved.
//
// This software is licensed under a Modified (3-Clause) BSD License.
// You should have received a copy of this license along with this
// software. If not, contact the copyright holder above.

// This code and comment are based by C source from:
//   https://github.com/Terminus-IMRC/rpi_cam_previewer/blob/f09489f164de3ece95f5e6c222995798d8bc379a/rpi_cam_previewer.c
// Thanks @Terminus-IMRC for publish simple camera preview source.

package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"unsafe"

	mmal "github.com/zchee/go-mmal"
	"github.com/zchee/go-mmal/bcmhost"
)

var (
	previewDestRect = mmal.NewRect(1024/4, 768/4, 1024/2, 768/4)
)

var (
	wait = flag.Int("time", 5000, "time for preview")
)

func main() {
	flag.Parse()

	bcmhost.Init()

	// create camera and video_render component
	var cpCamera mmal.WrapperType
	if err := mmal.WrapperCreate(&cpCamera, mmal.ComponentDefaultCamera); err != mmal.Success {
		log.Fatalf("go-picamera: %v", err)
	}
	var cpVideoRenderor mmal.WrapperType
	if err := mmal.WrapperCreate(&cpVideoRenderor, mmal.ComponentDefaultVideoRendereu); err != mmal.Success {
		log.Fatalf("go-picamera: %v", err)
	}

	// preview port of camera
	output := cpCamera.Output()
	mmal.SetPortFormat(&output)

	// input port of video_renderer
	input := cpVideoRenderor.Input()
	drhdr := mmal.NewParameterHeader(uint32(mmal.ParameterDisplayregion), uint32(unsafe.Sizeof(mmal.Displayregion{})))
	dr := mmal.NewDisplayregion(&drhdr, uint32(mmal.DisplaySetDestRect|mmal.DisplaySetFullscreen), false, previewDestRect)
	if err := mmal.PortParameterSet(&input, dr.Hdr()); err != mmal.Success {
		log.Fatalf("go-picamera: %v", err)
	}

	// create connection between camera's preview and video_renderer
	var connection mmal.ConnectionType
	cameraOutput := cpCamera.Output()
	videoInput := cpVideoRenderor.Input()
	if err := mmal.ConnectionCreate(&connection, &cameraOutput, &videoInput, mmal.ConnectionFlagTunnelling|mmal.ConnectionFlagAllocationOnInput); err != mmal.Success {
		log.Fatalf("go-picamera: %v", err)
	}
	// and enable it
	if err := mmal.ConnectionEnable(&connection); err != mmal.Success {
		log.Fatalf("go-picamera: %v", err)
	}

	destroy := func() {
		mmal.ConnectionDestroy(&connection)
		mmal.WrapperDestroy(&cpVideoRenderor)
		// TODO(zchee): occur "double free or corruption" error if destroy. Because already cleanuped from Go garbage collection?
		// mmal.WrapperDestroy(&cpCamera)
		bcmhost.Deinit()
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		destroy()
		os.Exit(1)
	}()

	// preview for "wait" seconds
	mmal.VcosSleep(uint32(*wait))
	destroy()
}
