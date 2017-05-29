// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/core/mmal_core_private.h>
#include <interface/mmal/mmal_logging.h>

extern void log_set_level(VCOS_LOG_LEVEL_T level) {
  vcos_log_set_level(VCOS_LOG_CATEGORY, level);
}
*/
import "C"

type LogLevel C.VCOS_LOG_LEVEL_T

const (
	LogLevelError LogLevel = C.VCOS_LOG_ERROR
	LogLevelInfo  LogLevel = C.VCOS_LOG_INFO
	LogLevelWarn  LogLevel = C.VCOS_LOG_WARN
	LogLevelDebug LogLevel = C.VCOS_LOG_INFO
	LogLevelTrace LogLevel = C.VCOS_LOG_TRACE
)

func LoggingInit() {
	C.mmal_logging_init()
}

func LoggingDeinit() {
	C.mmal_logging_deinit()
}

func LoggingSetLevel(level LogLevel) {
	C.log_set_level(C.VCOS_LOG_LEVEL_T(level))
}
