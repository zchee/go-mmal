// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_parameters_video.h>
*/
import "C"
import (
	"unsafe"
)

const (
	ParameterDisplayregion                   ParameterGroupVideo = (2 << 16) + iota // Takes a @ref MMAL_DISPLAYREGION_T
	ParameterSupportedProfiles                                                      // Takes a @ref MMAL_PARAMETER_VIDEO_PROFILE_T
	ParameterProfile                                                                // Takes a @ref MMAL_PARAMETER_VIDEO_PROFILE_T
	ParameterIntraperiod                                                            // Takes a @ref MMAL_PARAMETER_UINT32_T
	ParameterRatecontrol                                                            // Takes a @ref MMAL_PARAMETER_VIDEO_RATECONTROL_T
	ParameterNalunitformat                                                          // Takes a @ref MMAL_PARAMETER_VIDEO_NALUNITFORMAT_T
	ParameterMinimiseFragmentation                                                  // Takes a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterMbRowsPerSlice                                                         // Takes a @ref MMAL_PARAMETER_UINT32_T. Setting the value to zero resets to the default (one slice per frame)
	ParameterVideoLevelExtension                                                    // Takes a @ref MMAL_PARAMETER_VIDEO_LEVEL_EXTENSION_T
	ParameterVideoEedeEnable                                                        // Takes a @ref MMAL_PARAMETER_VIDEO_EEDE_ENABLE_T
	ParameterVideoEedeLossrate                                                      // Takes a @ref MMAL_PARAMETER_VIDEO_EEDE_LOSSRATE_T
	ParameterVideoRequestIFrame                                                     // Takes a @ref MMAL_PARAMETER_BOOLEAN_T. Request an I-frame
	ParameterVideoIntraRefresh                                                      // Takes a @ref MMAL_PARAMETER_VIDEO_INTRA_REFRESH_T
	ParameterVideoImmutableInput                                                    // Takes a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoBitRate                                                           // Takes a @ref MMAL_PARAMETER_UINT32_T. Run-time bit rate control
	ParameterVideoFrameRate                                                         // Takes a @ref MMAL_PARAMETER_FRAME_RATE_T
	ParameterVideoEncodeMinQuant                                                    // Takes a @ref MMAL_PARAMETER_UINT32_T
	ParameterVideoEncodeMaxQuant                                                    // Takes a @ref MMAL_PARAMETER_UINT32_T
	ParameterVideoEncodeRcModel                                                     // Takes a @ref MMAL_PARAMETER_VIDEO_ENCODE_RC_MODEL_T
	ParameterExtraBuffers                                                           // Takes a @ref MMAL_PARAMETER_UINT32_T
	ParameterVideoAlignHoriz                                                        // Takes a @ref MMAL_PARAMETER_UINT32_T. Changing this paramater from the default can reduce frame rate because image buffers need to be re-pitched.*/
	ParameterVideoAlignVert                                                         // Takes a @ref MMAL_PARAMETER_UINT32_T. Changing this paramater from the default can reduce frame rate because image buffers need to be re-pitched.*/
	ParameterVideoDroppablePframes                                                  // Take a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoEncodeInitialQuant                                                // Takes a @ref MMAL_PARAMETER_UINT32_T
	ParameterVideoEncodeQpP                                                         // Takes a @ref MMAL_PARAMETER_UINT32_T
	ParameterVideoEncodeRcSliceDquant                                               // Takes a @ref MMAL_PARAMETER_UINT32_T
	ParameterVideoEncodeFrameLimitBits                                              // Takes a @ref MMAL_PARAMETER_UINT32_T
	ParameterVideoEncodePeakRate                                                    // Takes a @ref MMAL_PARAMETER_UINT32_T
	ParameterVideoEncodeH264DisableCabac                                            // Take a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoEncodeH264LowLatency                                              // Take a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoEncodeH264AuDelimiters                                            // Take a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoEncodeH264DeblockIDC                                              // Takes a @ref MMAL_PARAMETER_UINT32_T
	ParameterVideoEncodeH264MbIntraMode                                             // Takes a @ref MMAL_PARAMETER_VIDEO_ENCODER_H264_MB_INTRA_MODES_T
	ParameterVideoEncodeHeaderOnOpen                                                // Takes a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoEncodePrecodeForQp                                                // Takes a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoDRMInitInfo                                                       // Takes a @ref MMAL_PARAMETER_VIDEO_DRM_INIT_INFO_T
	ParameterVideoTimestampFifo                                                     // Takes a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoDecodeErrorConcealment                                            // Takes a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoDRMProtectBuffer                                                  // Takes a @ref MMAL_PARAMETER_VIDEO_DRM_PROTECT_BUFFER_T
	ParameterVideoDecodeConfigVD3                                                   // Takes a @ref MMAL_PARAMETER_BYTES_T
	ParameterVideoEncodeH264VCLHRDParameters                                        // Take a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoEncodeH264LowDelayHRDFlag                                         // Take a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoEncodeInLineHeader                                                // Take a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoEncodeSeiEnable                                                   // Take a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoEncodeInlineVectors                                               // Take a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoRenderStats                                                       // Take a @ref MMAL_PARAMETER_VIDEO_RENDER_STATS_T
	ParameterVideoInterlaceType                                                     // Take a @ref MMAL_PARAMETER_VIDEO_INTERLACE_TYPE_T
	ParameterVideoInterpolateTimestamps                                             // Takes a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoEncodeSPSTiming                                                   // Take a @ref MMAL_PARAMETER_BOOLEAN_T
	ParameterVideoMaxNumCallbacks                                                   // Take a @ref MMAL_PARAMETER_UINT32_T
	ParameterVideoSourcePattern                                                     // Take a @ref MMAL_PARAMETER_SOURCE_PATTERN_T
)

type Displaytransform C.MMAL_DISPLAYTRANSFORM_T

const (
	DisplayRot0 Displaytransform = iota
	DisplayMirrorRot0
	DisplayMirrorRot180
	DisplayRot180
	DisplayMirrorRot90
	DisplayRot270
	DisplayRot90
	DisplayMirrorRot270
	DisplayDummy Displaytransform = 0x7FFFFFFF
)

type Displaymode C.MMAL_DISPLAYMODE_T

const (
	DisplayModeFill      Displaymode = iota
	DisplayModeLetterbox             // these allow a left eye source->dest to be specified and the right eye mapping will be inferred by symmetry
	DisplayModeStereoLeftToLeft
	DisplayModeStereoTopToTop
	DisplayModeStereoLeftToTop
	DisplayModeStereoTopToLeft
	DisplayModeDummy Displaymode = 0x7FFFFFFF
)

type Displayset C.MMAL_DISPLAYSET_T

const (
	DisplaySetNone        Displayset = 0
	DisplaySetNum         Displayset = 1
	DisplaySetFullscreen  Displayset = 2
	DisplaySetTransform   Displayset = 4
	DisplaySetDestRect    Displayset = 8
	DisplaySetSrcRect     Displayset = 0x10
	DisplaySetMode        Displayset = 0x20
	DisplaySetPixel       Displayset = 0x40
	DisplaySetNoaspect    Displayset = 0x80
	DisplaySetLayer       Displayset = 0x100
	DisplaySetCopyprotect Displayset = 0x200
	DisplaySetAlpha       Displayset = 0x400
	DisplaySetDummy       Displayset = 0x7FFFFFFF
)

type Displayregion struct {
	c C.MMAL_DISPLAYREGION_T
}

func NewDisplayregion(header *ParameterHeader, set uint32, fullscreen bool, previewDestRect Rect) Displayregion {
	var c C.MMAL_DISPLAYREGION_T
	c.hdr = header.c
	c.set = C.uint32_t(set)
	c.fullscreen = 0
	if fullscreen {
		c.fullscreen = 1
	}
	c.dest_rect = previewDestRect.c
	return Displayregion{c}
}

func (d Displayregion) Hdr() ParameterHeader {
	return ParameterHeader{d.c.hdr}
}

func (d Displayregion) Set() uint32 {
	return uint32(d.c.set)
}

func (d Displayregion) DisplayNum() uint32 {
	return uint32(d.c.display_num)
}

func (d Displayregion) Fullscreen() bool {
	return BoolT(d.c.fullscreen) != BoolT(0)
}

func (d Displayregion) Transform() Displaytransform {
	return Displaytransform(d.c.transform)
}

func (d Displayregion) DestRect() Rect {
	return Rect{d.c.dest_rect}
}

func (d Displayregion) SrcRect() Rect {
	return Rect{d.c.src_rect}
}

func (d Displayregion) Noaspect() bool {
	return BoolT(d.c.noaspect) != BoolT(0)
}

func (d Displayregion) Mode() Displaymode {
	return Displaymode(d.c.mode)
}

func (d Displayregion) PixelX() uint32 {
	return uint32(d.c.pixel_x)
}

func (d Displayregion) PixelY() uint32 {
	return uint32(d.c.pixel_y)
}

func (d Displayregion) Layer() int32 {
	return int32(d.c.layer)
}

func (d Displayregion) CopyprotectRequired() bool {
	return BoolT(d.c.copyprotect_required) != BoolT(0)
}

func (d Displayregion) Alpha() uint32 {
	return uint32(d.c.alpha)
}

type VideoProfile C.MMAL_VIDEO_PROFILE_T

const (
	VideoProfileH263Baseline VideoProfile = iota
	VideoProfileH263H320coding
	VideoProfileH263Backwardcompatible
	VideoProfileH263Iswv2
	VideoProfileH263Iswv3
	VideoProfileH263Highcompression
	VideoProfileH263Internet
	VideoProfileH263Interlace
	VideoProfileH263Highlatency
	VideoProfileMP4VSimple
	VideoProfileMP4VSimplescalable
	VideoProfileMP4VCore
	VideoProfileMP4VMain
	VideoProfileMP4VNbit
	VideoProfileMP4VScalabletexture
	VideoProfileMP4VSimpleface
	VideoProfileMP4VSimplefba
	VideoProfileMP4VBasicanimated
	VideoProfileMP4VHybrid
	VideoProfileMP4VAdvancedrealtime
	VideoProfileMP4VCorescalable
	VideoProfileMP4VAdvancedcoding
	VideoProfileMP4VAdvancedcore
	VideoProfileMP4VAdvancedscalable
	VideoProfileMP4VAdvancedsimple
	VideoProfileH264Baseline
	VideoProfileH264Main
	VideoProfileH264Extended
	VideoProfileH264High
	VideoProfileH264High10
	VideoProfileH264High422
	VideoProfileH264High444
	VideoProfileH264Constrained_baseline
	VideoProfileDummy VideoProfile = 0x7FFFFFFF
)

type VideoLevel C.MMAL_VIDEO_LEVEL_T

const (
	VideoLevelH26310 VideoLevel = iota
	VideoLevelH26320
	VideoLevelH26330
	VideoLevelH26340
	VideoLevelH26345
	VideoLevelH26350
	VideoLevelH26360
	VideoLevelH26370
	VideoLevelMP4V0
	VideoLevelMP4V0b
	VideoLevelMP4V1
	VideoLevelMP4V2
	VideoLevelMP4V3
	VideoLevelMP4V4
	VideoLevelMP4V4a
	VideoLevelMP4V5
	VideoLevelMP4V6
	VideoLevelH2641
	VideoLevelH2641b
	VideoLevelH26411
	VideoLevelH26412
	VideoLevelH26413
	VideoLevelH2642
	VideoLevelH26421
	VideoLevelH26422
	VideoLevelH2643
	VideoLevelH26431
	VideoLevelH26432
	VideoLevelH2644
	VideoLevelH26441
	VideoLevelH26442
	VideoLevelH2645
	VideoLevelH26451
	VideoLevelDummy VideoLevel = 0x7FFFFFFF
)

type ParameterVideoProfile struct {
	c C.MMAL_PARAMETER_VIDEO_PROFILE_T
}

func (p ParameterVideoProfile) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

type ParameterVideoProfileProfile struct {
	profile VideoProfile
	level   VideoLevel
}

func (p ParameterVideoProfileProfile) Profile() VideoProfile {
	return p.profile
}

func (p ParameterVideoProfileProfile) Level() VideoLevel {
	return p.level
}

func (p ParameterVideoProfile) Profile() ParameterVideoProfileProfile {
	return ParameterVideoProfileProfile{
		profile: VideoProfile(p.c.profile[0].profile),
		level:   VideoLevel(p.c.profile[0].level),
	}
}

type VideoRatecontrol C.MMAL_VIDEO_RATECONTROL_T

const (
	VideoRatecontrolDefault VideoRatecontrol = iota
	VideoRatecontrolVariable
	VideoRatecontrolConstant
	VideoRatecontrolVariableSkipFrames
	VideoRatecontrolConstantSkipFrames
	VideoRatecontrolDummy VideoRatecontrol = 0x7fffffff
)

type VideoIntraRefresh C.MMAL_VIDEO_INTRA_REFRESH_T

const (
	VideoIntraRefreshCyclic VideoIntraRefresh = iota
	VideoIntraRefreshAdaptive
	VideoIntraRefreshBoth
	VideoIntraRefreshKhronosextensions VideoIntraRefresh = 0x6F000000
	VideoIntraRefreshVendorstartunused VideoIntraRefresh = 0x7F000000
	VideoIntraRefreshCyclicMrows
	VideoIntraRefreshPseudoRand
	VideoIntraRefreshMax
	VideoIntraRefreshDummy VideoIntraRefresh = 0x7FFFFFFF
)

type VideoEncodeRCModel C.MMAL_VIDEO_ENCODE_RC_MODEL_T

const (
	VideoEncodeRCModelDefault VideoEncodeRCModel = iota
	VideoEncodeRCModelJVT                        = VideoEncodeRCModelDefault
	VideoEncodeRCModelVOWIFI
	VideoEncodeRCModelCBR
	VideoEncodeRCModelLast
	VideoEncodeRCModelDummy VideoEncodeRCModel = 0x7FFFFFFF
)

type ParameterVideoEncodeRCModel struct {
	c C.MMAL_PARAMETER_VIDEO_ENCODE_RC_MODEL_T
}

func (p ParameterVideoEncodeRCModel) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterVideoEncodeRCModel) RCModel() VideoEncodeRCModel {
	return VideoEncodeRCModel(p.c.rc_model)
}

type ParameterVideoEncodeRatecontrol struct {
	c C.MMAL_PARAMETER_VIDEO_ENCODE_RC_MODEL_T
}

func (p ParameterVideoEncodeRatecontrol) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterVideoEncodeRatecontrol) Control() VideoRatecontrol {
	return VideoRatecontrol(p.c.rc_model)
}

type VideoEncodeH264MBIntraModes C.MMAL_VIDEO_ENCODE_H264_MB_INTRA_MODES_T

const (
	MMAL_VIDEO_ENCODER_H264_MB_4x4_INTRA VideoEncodeH264MBIntraModes = 1 << iota
	MMAL_VIDEO_ENCODER_H264_MB_8x8_INTRA
	MMAL_VIDEO_ENCODER_H264_MB_16x16_INTRA
	MMAL_VIDEO_ENCODER_H264_MB_INTRA_DUMMY VideoEncodeH264MBIntraModes = 0x7ffffff
)

type ParameterVideoEncoderH264MBIntraModes struct {
	c C.MMAL_PARAMETER_VIDEO_ENCODER_H264_MB_INTRA_MODES_T
}

func (p ParameterVideoEncoderH264MBIntraModes) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterVideoEncoderH264MBIntraModes) MBMode() VideoEncodeH264MBIntraModes {
	return VideoEncodeH264MBIntraModes(p.c.mb_mode)
}

type VideoNalunitformat C.MMAL_VIDEO_NALUNITFORMAT_T

const (
	VideoNalunitformatStartcodes VideoNalunitformat = 1 << iota
	VideoNalunitformatNalunitperbuffer
	VideoNalunitformatOnebyteinterleavelength
	VideoNalunitformatTwobyteinterleavelength
	VideoNalunitformatFourbyteinterleavelength
	VideoNalunitformatDummy VideoNalunitformat = 0x7fffffff
)

type ParameterVideoNalunitformat struct {
	c C.MMAL_PARAMETER_VIDEO_NALUNITFORMAT_T
}

func (p ParameterVideoNalunitformat) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterVideoNalunitformat) Format() VideoNalunitformat {
	return VideoNalunitformat(p.c.format)
}

type ParameterVideoLevelExtensionType struct {
	c C.MMAL_PARAMETER_VIDEO_LEVEL_EXTENSION_T
}

func (p ParameterVideoLevelExtensionType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterVideoLevelExtensionType) CustomMaxMBPS() uint32 {
	return uint32(p.c.custom_max_mbps)
}

func (p ParameterVideoLevelExtensionType) CustomMaxFS() uint32 {
	return uint32(p.c.custom_max_fs)
}

func (p ParameterVideoLevelExtensionType) CustomMaxBrAndCpb() uint32 {
	return uint32(p.c.custom_max_br_and_cpb)
}

type ParameterVideoIntraRefreshType struct {
	c C.MMAL_PARAMETER_VIDEO_INTRA_REFRESH_T
}

func (p ParameterVideoIntraRefreshType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterVideoIntraRefreshType) RefreshMode() VideoIntraRefresh {
	return VideoIntraRefresh(p.c.refresh_mode)
}

func (p ParameterVideoIntraRefreshType) AirMBS() uint32 {
	return uint32(p.c.air_mbs)
}

func (p ParameterVideoIntraRefreshType) AirRef() uint32 {
	return uint32(p.c.air_ref)
}

func (p ParameterVideoIntraRefreshType) CirMBS() uint32 {
	return uint32(p.c.cir_mbs)
}

func (p ParameterVideoIntraRefreshType) PirMBS() uint32 {
	return uint32(p.c.pir_mbs)
}

type ParameterVideoEedeEnableType struct {
	c C.MMAL_PARAMETER_VIDEO_EEDE_ENABLE_T
}

func (p ParameterVideoEedeEnableType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterVideoEedeEnableType) Enable() uint32 {
	return uint32(p.c.enable)
}

type ParameterVideoEedeLossrateType struct {
	c C.MMAL_PARAMETER_VIDEO_EEDE_LOSSRATE_T
}

func (p ParameterVideoEedeLossrateType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterVideoEedeLossrateType) LossRate() uint32 {
	return uint32(p.c.loss_rate)
}

type ParameterVideoDRMInitInfoType struct {
	c C.MMAL_PARAMETER_VIDEO_DRM_INIT_INFO_T
}

func (p ParameterVideoDRMInitInfoType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterVideoDRMInitInfoType) CurrentTime() uint32 {
	return uint32(p.c.current_time)
}

func (p ParameterVideoDRMInitInfoType) TicksPerSec() uint32 {
	return uint32(p.c.ticks_per_sec)
}

func (p ParameterVideoDRMInitInfoType) Lhs() [32]uint8 {
	lhses := [32]uint8{}
	for i, lhs := range p.c.lhs {
		lhses[i] = uint8(lhs)
	}
	return lhses
}

type ParameterVideoDRMProtectBufferType struct {
	c C.MMAL_PARAMETER_VIDEO_DRM_PROTECT_BUFFER_T
}

func (p ParameterVideoDRMProtectBufferType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterVideoDRMProtectBufferType) SizeWanted() uint32 {
	return uint32(p.c.size_wanted)
}

func (p ParameterVideoDRMProtectBufferType) Protect() uint32 {
	return uint32(p.c.protect)
}

func (p ParameterVideoDRMProtectBufferType) MemHandle() uint32 {
	return uint32(p.c.mem_handle)
}

func (p ParameterVideoDRMProtectBufferType) PhysAddr() unsafe.Pointer {
	return p.c.phys_addr
}

type ParameterVideoRenderStatsType struct {
	c C.MMAL_PARAMETER_VIDEO_RENDER_STATS_T
}

func (p ParameterVideoRenderStatsType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterVideoRenderStatsType) Valid() bool {
	return BoolT(p.c.valid) != BoolT(0)
}

func (p ParameterVideoRenderStatsType) Match() uint32 {
	return uint32(p.c.match)
}

func (p ParameterVideoRenderStatsType) Period() uint32 {
	return uint32(p.c.period)
}

func (p ParameterVideoRenderStatsType) Phase() uint32 {
	return uint32(p.c.phase)
}

func (p ParameterVideoRenderStatsType) PixelClockNominal() uint32 {
	return uint32(p.c.pixel_clock_nominal)
}

func (p ParameterVideoRenderStatsType) PixelClock() uint32 {
	return uint32(p.c.pixel_clock)
}

func (p ParameterVideoRenderStatsType) HVSStatus() uint32 {
	return uint32(p.c.hvs_status)
}

func (p ParameterVideoRenderStatsType) Dummy() [2]uint32 {
	dummies := [2]uint32{}
	for i, dummy := range p.c.dummy {
		dummies[i] = uint32(dummy)
	}
	return dummies
}

type Interlacetype C.MMAL_INTERLACETYPE_T

const (
	InterlaceProgressive                 Interlacetype = iota       // The data is not interlaced, it is progressive scan
	InterlaceFieldSingleUpperFirst                                  // The data is interlaced, fields sent separately in temporal order, with upper field first
	InterlaceFieldSingleLowerFirst                                  // The data is interlaced, fields sent separately in temporal order, with lower field first
	InterlaceFieldsInterleavedUpperFirst                            // The data is interlaced, two fields sent together line interleaved, with the upper field temporally earlier
	InterlaceFieldsInterleavedLowerFirst                            // The data is interlaced, two fields sent together line interleaved, with the lower field temporally earlier
	InterlaceMixed                                                  // The stream may contain a mixture of progressive and interlaced frames
	InterlaceKhronosExtensions           Interlacetype = 0x6F000000 // Reserved region for introducing Khronos Standard Extensions
	InterlaceVendorStartUnused           Interlacetype = 0x7F000000 // Reserved region for introducing Vendor Extensions
	InterlaceMax                         Interlacetype = 0x7FFFFFFF
)

type ParameterVideoInterlaceTypeType struct {
	c C.MMAL_PARAMETER_VIDEO_INTERLACE_TYPE_T
}

func (p ParameterVideoInterlaceTypeType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterVideoInterlaceTypeType) EMode() Interlacetype {
	return Interlacetype(p.c.eMode)
}

func (p ParameterVideoInterlaceTypeType) BRepeatFirstField() bool {
	return BoolT(p.c.bRepeatFirstField) != BoolT(0)
}

type SourcePattern C.MMAL_SOURCE_PATTERN_T

const (
	VideoSourcePatternWhite SourcePattern = iota
	VideoSourcePatternBlack
	VideoSourcePatternDiagonal
	VideoSourcePatternNoise
	VideoSourcePatternRandom
	VideoSourcePatternColour
	VideoSourcePatternBlocks
	VideoSourcePatternSwirly
	VideoSourcePatternDummy SourcePattern = 0x7ffffff
)

type ParameterVideoSourcePatternType struct {
	c C.MMAL_PARAMETER_VIDEO_SOURCE_PATTERN_T
}

func (p ParameterVideoSourcePatternType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterVideoSourcePatternType) Pattern() SourcePattern {
	return SourcePattern(p.c.pattern)
}
