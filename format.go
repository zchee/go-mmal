// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal_format.h>
*/
import "C"

type EsType int

const (
	EsTypeUnKnown EsType = iota
	EsTypeControl
	EsTypeAudio
	EsTypeVideo
	EsTypeSubpicture
)

type VideoFormat struct {
	c C.MMAL_VIDEO_FORMAT_T
}

func (v VideoFormat) Width() uint32 {
	return uint32(v.c.width)
}

func (v VideoFormat) Height() uint32 {
	return uint32(v.c.height)
}

func (v VideoFormat) Crop() Rect {
	return Rect{v.c.crop}
}

func (v VideoFormat) FrameRate() Rational {
	return Rational{v.c.frame_rate}
}

func (v VideoFormat) Par() Rational {
	return Rational{v.c.par}
}

func (v VideoFormat) ColorSpace() FourCCT {
	return FourCCT(v.c.color_space)
}

type AudioFormat struct {
	c C.MMAL_AUDIO_FORMAT_T
}

func (a AudioFormat) Channels() uint32 {
	return uint32(a.c.channels)
}
func (a AudioFormat) SampleRate() uint32 {
	return uint32(a.c.sample_rate)
}

func (a AudioFormat) BitsPerSample() uint32 {
	return uint32(a.c.bits_per_sample)
}

func (a AudioFormat) BlockAlign() uint32 {
	return uint32(a.c.block_align)
}

type SubpictureFormat struct {
	c C.MMAL_SUBPICTURE_FORMAT_T
}

func (s SubpictureFormat) XOffset() uint32 {
	return uint32(s.c.x_offset)
}

func (s SubpictureFormat) YOffset() uint32 {
	return uint32(s.c.y_offset)
}

// TODO(zchee): union type
type EsSpecificFormat struct {
	c *C.MMAL_ES_SPECIFIC_FORMAT_T
}

const EsFormatFlagFramed = 0x1

const EncodingVariantDefault = 0

type EsFormat struct {
	c *C.MMAL_ES_FORMAT_T
}

func (e EsFormat) Type() EsType {
	return EsType(e.c._type)
}

func (e EsFormat) Encoding() FourCCT {
	return FourCCT(e.c.encoding)
}

func (e EsFormat) EncodingVariant() FourCCT {
	return FourCCT(e.c.encoding_variant)
}

func (e EsFormat) ES() EsSpecificFormat {
	return EsSpecificFormat{e.c.es}
}

func (e EsFormat) Bitrate() uint32 {
	return uint32(e.c.bitrate)
}

func (e EsFormat) Flags() uint32 {
	return uint32(e.c.flags)
}

func (e EsFormat) ExtradataSize() uint32 {
	return uint32(e.c.extradata_size)
}

func (e EsFormat) Extradata() uint8 {
	return uint8(*e.c.extradata)
}

func FormatAlloc() EsFormat {
	return EsFormat{
		c: C.mmal_format_alloc(),
	}
}

func (e EsFormat) Free() {
	C.mmal_format_free(e.c)
}

func FormatExtradatAlloc(format EsFormat, size uint) Status {
	return Status(C.mmal_format_extradata_alloc(format.c, C.uint(size)))
}

func FormatCopy(formatDest, formatSrc EsFormat) {
	C.mmal_format_copy(formatDest.c, formatSrc.c)
}

func FormatFullCopy(formatDest, formatSrc EsFormat) {
	C.mmal_format_full_copy(formatDest.c, formatSrc.c)
}

const (
	EsFormatCompareFlagType      = 0x01 /**< The type is different */
	EsFormatCompareFlagEncoding  = 0x02 /**< The encoding is different */
	EsFormatCompareFlagBitrate   = 0x04 /**< The bitrate is different */
	EsFormatCompareFlagFlags     = 0x08 /**< The flags are different */
	EsFormatCompareFlagExtradata = 0x10 /**< The extradata is different */

	EsFormatCompareFlagVideoResolution   = 0x0100 /**< The video resolution is different */
	EsFormatCompareFlagVideoCropping     = 0x0200 /**< The video cropping is different */
	EsFormatCompareFlagVideoFrame_rate   = 0x0400 /**< The video frame rate is different */
	EsFormatCompareFlagVideoAspect_ratio = 0x0800 /**< The video aspect ratio is different */
	EsFormatCompareFlagVideoColor_space  = 0x1000 /**< The video color space is different */

	EsFormatCompareFlagEsOther = 0x10000000 /**< Other ES specific parameters are different */
)

func FormatCompare(format1, format2 EsFormat) C.uint32_t {
	return C.mmal_format_compare(format1.c, format2.c)
}
