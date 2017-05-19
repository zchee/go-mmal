// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_parameters_common.h>
*/
import "C"

// Common parameter ID group, used with many types of component.
type ParameterGroupCommon int // (0 << 16)
// Camera-specific parameter ID group.
type ParameterGroupCamera int // (1 << 16)
// Video-specific parameter ID group.
type ParameterGroupVideo int // (2 << 16)
// Audio-specific parameter ID group.
type ParameterGroupAudio int // (3 << 16)
// Clock-specific parameter ID group.
type ParameterGroupClock int // (4 << 16)
// Miracast-specific parameter ID group.
type ParameterGroupMiracast int // (5 << 16)

// Common MMAL parameter IDs.
const (
	ParameterUnused             ParameterGroupCommon = (0 << 16) + iota // < Never a valid parameter ID
	ParameterSupportedEncodings                                         // < Takes a MMAL_PARAMETER_ENCODING_T
	ParameterURI                                                        // < Takes a MMAL_PARAMETER_URI_T
	ParameterChangeEventRequest                                         // < Takes a MMAL_PARAMETER_CHANGE_EVENT_REQUEST_T
	ParameterZeroCopy                                                   // < Takes a MMAL_PARAMETER_BOOLEAN_T
	ParameterBufferRequirements                                         // < Takes a MMAL_PARAMETER_BUFFER_REQUIREMENTS_T
	ParameterStatistics                                                 // < Takes a MMAL_PARAMETER_STATISTICS_T
	ParameterCoreStatistics                                             // < Takes a MMAL_PARAMETER_CORE_STATISTICS_T
	ParameterMemUsage                                                   // < Takes a MMAL_PARAMETER_MEM_USAGE_T
	ParameterBufferFlagFilter                                           // < Takes a MMAL_PARAMETER_UINT32_T
	ParameterSeek                                                       // < Takes a MMAL_PARAMETER_SEEK_T
	ParameterPowermonEnable                                             // < Takes a MMAL_PARAMETER_BOOLEAN_T
	ParameterLogging                                                    // < Takes a MMAL_PARAMETER_LOGGING_T
	ParameterSystemtime                                                 // < Takes a MMAL_PARAMETER_UINT64_T
	ParameterNoImagePadding                                             // < Takes a MMAL_PARAMETER_BOOLEAN_T
	ParameterLockstepEnable                                             // < Takes a MMAL_PARAMETER_BOOLEAN_T
)

type ParameterHeader struct {
	c C.MMAL_PARAMETER_HEADER_T
}

func (p ParameterHeader) ID() uint32 {
	return uint32(p.c.id)
}

func (p ParameterHeader) Size() uint32 {
	return uint32(p.c.size)
}

type ParameterChangeEventRequestType struct {
	c C.MMAL_PARAMETER_CHANGE_EVENT_REQUEST_T
}

func (p ParameterChangeEventRequestType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterChangeEventRequestType) ChangeID() uint32 {
	return uint32(p.c.change_id)
}

func (p ParameterChangeEventRequestType) Enable() BoolT {
	return BoolT(p.c.enable)
}

type ParameterBufferRequirementsType struct {
	c C.MMAL_PARAMETER_BUFFER_REQUIREMENTS_T
}

func (p ParameterBufferRequirementsType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterBufferRequirementsType) BufferNumMin() uint32 {
	return uint32(p.c.buffer_num_min)
}

func (p ParameterBufferRequirementsType) BufferSizeMin() uint32 {
	return uint32(p.c.buffer_size_min)
}

func (p ParameterBufferRequirementsType) BufferAlignmentMin() uint32 {
	return uint32(p.c.buffer_alignment_min)
}

func (p ParameterBufferRequirementsType) BufferNumRecommended() uint32 {
	return uint32(p.c.buffer_num_recommended)
}

func (p ParameterBufferRequirementsType) BufferSizeRecommended() uint32 {
	return uint32(p.c.buffer_size_recommended)
}

type ParameterSeekType struct {
	c C.MMAL_PARAMETER_SEEK_T
}

func (p ParameterSeekType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterSeekType) Offset() int64 {
	return int64(p.c.offset)
}

func (p ParameterSeekType) Flags() uint32 {
	return uint32(p.c.flags)
}

const (
	ParamSeekFlagPrecise = 0x1 // < Choose precise seeking even if slower
	ParamSeekFlagForward = 0x2 // < Seek to the next keyframe following the specified offset
)

type ParameterStatisticsType struct {
	c C.MMAL_PARAMETER_STATISTICS_T
}

func (p ParameterStatisticsType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterStatisticsType) BufferCount() uint32 {
	return uint32(p.c.buffer_count)
}

func (p ParameterStatisticsType) FrameCount() uint32 {
	return uint32(p.c.frame_count)
}

func (p ParameterStatisticsType) FramesSkipped() uint32 {
	return uint32(p.c.frames_skipped)
}

func (p ParameterStatisticsType) FramesDiscarded() uint32 {
	return uint32(p.c.frames_discarded)
}

func (p ParameterStatisticsType) EOSSeen() uint32 {
	return uint32(p.c.eos_seen)
}

func (p ParameterStatisticsType) MaximumFrameBytes() uint32 {
	return uint32(p.c.maximum_frame_bytes)
}

func (p ParameterStatisticsType) TotalBytes() int64 {
	return int64(p.c.total_bytes)
}

func (p ParameterStatisticsType) CorruptMacroblocks() uint32 {
	return uint32(p.c.corrupt_macroblocks)
}

type CoreStatsDir C.MMAL_CORE_STATS_DIR

const (
	CoreStatsRX CoreStatsDir = iota
	CoreStatsTX
	CoreStatsMax CoreStatsDir = 0x7fffffff
)

type ParameterCoreStatisticsType struct {
	c C.MMAL_PARAMETER_CORE_STATISTICS_T
}

func (p ParameterCoreStatisticsType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCoreStatisticsType) dir() CoreStatsDir {
	return CoreStatsDir(p.c.dir)
}

func (p ParameterCoreStatisticsType) Reset() BoolT {
	return BoolT(p.c.reset)
}

// TODO(zchee): implements
// func (p ParameterCoreStatisticsType) Stats() CoreStatistics {}

type ParameterMemUsageType struct {
	c C.MMAL_PARAMETER_MEM_USAGE_T
}

func (p ParameterMemUsageType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterMemUsageType) PoolMemAllocSize() uint32 {
	return uint32(p.c.pool_mem_alloc_size)
}

type ParameterLoggingType struct {
	c C.MMAL_PARAMETER_LOGGING_T
}

func (p ParameterLoggingType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterLoggingType) Set() uint32 {
	return uint32(p.c.set)
}

func (p ParameterLoggingType) Clear() uint32 {
	return uint32(p.c.clear)
}
