// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build standalone

package mmal

const (
	ComponentDefaultVideoDecoder   = "avcodec.video_decode"
	ComponentDefaultVideoEncoder   = "avcodec.video_encode"
	ComponentDefaultVideoRendereu  = "sdl.video_render"
	ComponentDefaultImageDecoder   = "avcodec.video_decode"
	ComponentDefaultImageEncoder   = "avcodec.video_encode"
	ComponentDefaultCamera         = "artificial_camera"
	ComponentDefaultVideoConverter = "avcodec.video_convert"
	ComponentDefaultSplitter       = "splitter"
	ComponentDefaultScheduler      = "scheduler"
	ComponentDefaultVideoInjecter  = "video_inject"
	ComponentDefaultAudioDecoder   = "avcodec.audio_decode"
	ComponentDefaultAudioRenderer  = "sdl.audio_render"
	ComponentDefaultMiracast       = "miracast"
	ComponentDefaultClock          = "clock"
)
