// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_parameters.h>
*/
import "C"

type ParameterUint64Type struct {
	c C.MMAL_PARAMETER_UINT64_T
}

func (p ParameterUint64Type) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterUint64Type) Value() uint64 {
	return uint64(p.c.value)
}

type ParameterInt64Type struct {
	c C.MMAL_PARAMETER_INT64_T
}

func (p ParameterInt64Type) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterInt64Type) Value() int64 {
	return int64(p.c.value)
}

type ParameterUint32Type struct {
	c C.MMAL_PARAMETER_UINT32_T
}

func (p ParameterUint32Type) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterUint32Type) Value() uint32 {
	return uint32(p.c.value)
}

type ParameterInt32Type struct {
	c C.MMAL_PARAMETER_INT32_T
}

func (p ParameterInt32Type) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterInt32Type) Value() int32 {
	return int32(p.c.value)
}

type ParameterRational struct {
	c C.MMAL_PARAMETER_RATIONAL_T
}

func (p ParameterRational) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterRational) Value() Rational {
	return Rational{p.c.value}
}

type ParameterBooleanType struct {
	c C.MMAL_PARAMETER_BOOLEAN_T
}

func (p ParameterBooleanType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterBooleanType) Enable() bool {
	return BoolT(p.c.enable) != BoolT(0)
}

type ParameterStringType struct {
	c C.MMAL_PARAMETER_STRING_T
}

func (p ParameterStringType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterStringType) Str() string {
	return C.GoString(&p.c.str[0])
}

type ParameterBytesType struct {
	c C.MMAL_PARAMETER_BYTES_T
}

func (p ParameterBytesType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterBytesType) Data() uint8 {
	return uint8(p.c.data[0])
}

type ParameterScalefactorType struct {
	c C.MMAL_PARAMETER_SCALEFACTOR_T
}

func (p ParameterScalefactorType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterScalefactorType) ScaleX() Fixed1616 {
	return Fixed1616(p.c.scale_x)
}

func (p ParameterScalefactorType) ScaleY() Fixed1616 {
	return Fixed1616(p.c.scale_y)
}

type ParamMirrorType C.MMAL_PARAM_MIRROR_T

const (
	ParamMirrorNone ParamMirrorType = iota
	ParamMirrorVertical
	ParamMirrorHorizontal
	ParamMirrorBoth
)

type ParameterMirrorType struct {
	c C.MMAL_PARAMETER_MIRROR_T
}

func (p ParameterMirrorType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterMirrorType) Value() ParamMirrorType {
	return ParamMirrorType(p.c.value)
}

type ParameterUriType struct {
	c C.MMAL_PARAMETER_URI_T
}

func (p ParameterUriType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterUriType) URI() string {
	return C.GoString(&p.c.uri[0])
}

type ParameterEncodingType struct {
	c C.MMAL_PARAMETER_ENCODING_T
}

func (p ParameterEncodingType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterEncodingType) Encoding() uint32 {
	return uint32(p.c.encoding[0])
}

type ParameterFrameRateType struct {
	c C.MMAL_PARAMETER_FRAME_RATE_T
}

func (p ParameterFrameRateType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterFrameRateType) FrameRate() Rational {
	return Rational{p.c.frame_rate}
}

type ParameterConfigfileType struct {
	c C.MMAL_PARAMETER_CONFIGFILE_T
}

func (p ParameterConfigfileType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterConfigfileType) FileSize() uint32 {
	return uint32(p.c.file_size)
}

type ParameterConfigfileChunkType struct {
	c C.MMAL_PARAMETER_CONFIGFILE_CHUNK_T
}

func (p ParameterConfigfileChunkType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterConfigfileChunkType) Size() uint32 {
	return uint32(p.c.size)
}

func (p ParameterConfigfileChunkType) Offset() uint32 {
	return uint32(p.c.offset)
}

func (p ParameterConfigfileChunkType) Data() string {
	return C.GoString(&p.c.data[0])
}
