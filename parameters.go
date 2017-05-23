// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_parameters.h>
*/
import "C"
import (
	"unsafe"
)

type ParameterUint64Type struct {
	c C.MMAL_PARAMETER_UINT64_T
}

func NewParameterUint64Type(header ParameterHeader, value uint64) ParameterUint64Type {
	var c C.MMAL_PARAMETER_UINT64_T
	c.hdr = header.c
	c.value = C.uint64_t(value)
	return ParameterUint64Type{c}
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

func NewParameterInt64Type(header ParameterHeader, value int64) ParameterInt64Type {
	var c C.MMAL_PARAMETER_INT64_T
	c.hdr = header.c
	c.value = C.int64_t(value)
	return ParameterInt64Type{c}
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

func NewParameterUint32Type(header ParameterHeader, value uint64) ParameterUint32Type {
	var c C.MMAL_PARAMETER_UINT32_T
	c.hdr = header.c
	c.value = C.uint32_t(value)
	return ParameterUint32Type{c}
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

func NewParameterInt32Type(header ParameterHeader, value int32) ParameterInt32Type {
	var c C.MMAL_PARAMETER_INT32_T
	c.hdr = header.c
	c.value = C.int32_t(value)
	return ParameterInt32Type{c}
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

func NewParameterRational(header ParameterHeader, value Rational) ParameterRational {
	var c C.MMAL_PARAMETER_RATIONAL_T
	c.hdr = header.c
	c.value = value.c
	return ParameterRational{c}
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

func NewParameterBooleanType(header ParameterHeader, enable bool) ParameterBooleanType {
	var c C.MMAL_PARAMETER_BOOLEAN_T
	c.hdr = header.c
	isEnable := C.MMAL_BOOL_T(0)
	if enable {
		isEnable = C.MMAL_BOOL_T(1)
	}
	c.enable = isEnable
	return ParameterBooleanType{c}
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

func NewParameterStringType(header ParameterHeader, str string) ParameterStringType {
	var c C.MMAL_PARAMETER_STRING_T
	c.hdr = header.c
	var strsl [1]C.char
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	strsl[0] = *cStr
	c.str = strsl
	return ParameterStringType{c}
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

func NewParameterBytesType(header ParameterHeader, data uint8) ParameterBytesType {
	var c C.MMAL_PARAMETER_BYTES_T
	c.hdr = header.c
	var bytesl [1]C.uint8_t
	bytesl[0] = C.uint8_t(data)
	c.data = bytesl
	return ParameterBytesType{c}
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

func NewParameterScalefactorType(header ParameterHeader, scalex, scaley uint32) ParameterScalefactorType {
	var c C.MMAL_PARAMETER_SCALEFACTOR_T
	c.hdr = header.c
	c.scale_x = C.MMAL_FIXED_16_16_T(scalex)
	c.scale_y = C.MMAL_FIXED_16_16_T(scaley)
	return ParameterScalefactorType{c}
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

func NewParameterMirrorType(header ParameterHeader, value ParamMirrorType) ParameterMirrorType {
	var c C.MMAL_PARAMETER_MIRROR_T
	c.hdr = header.c
	c.value = C.MMAL_PARAM_MIRROR_T(value)
	return ParameterMirrorType{c}
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

func NewParameterUriType(header ParameterHeader, uri string) ParameterUriType {
	var c C.MMAL_PARAMETER_URI_T
	c.hdr = header.c
	var urisl [1]C.char
	u := C.CString(uri)
	defer C.free(unsafe.Pointer(u))
	urisl[0] = *u
	c.uri = urisl
	return ParameterUriType{c}
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

func NewParameterEncodingType(header ParameterHeader, encoding uint32) ParameterEncodingType {
	var c C.MMAL_PARAMETER_ENCODING_T
	c.hdr = header.c
	var encodingsl [1]C.uint32_t
	encodingsl[0] = C.uint32_t(encoding)
	c.encoding = encodingsl
	return ParameterEncodingType{c}
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

func NewParameterFrameRateType(header ParameterHeader, frameRate Rational) ParameterFrameRateType {
	var c C.MMAL_PARAMETER_FRAME_RATE_T
	c.hdr = header.c
	c.frame_rate = C.MMAL_RATIONAL_T(frameRate.c)
	return ParameterFrameRateType{c}
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

func NewParameterConfigfileType(header ParameterHeader, filesize uint32) ParameterConfigfileType {
	var c C.MMAL_PARAMETER_CONFIGFILE_T
	c.hdr = header.c
	c.file_size = C.uint32_t(filesize)
	return ParameterConfigfileType{c}
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

func NewParameterConfigfileChunkType(header ParameterHeader, size, offset uint32, data string) ParameterConfigfileChunkType {
	var c C.MMAL_PARAMETER_CONFIGFILE_CHUNK_T
	c.hdr = header.c
	c.size = C.uint32_t(size)
	c.offset = C.uint32_t(offset)
	var dsl [1]C.char
	d := C.CString(data)
	defer C.free(unsafe.Pointer(d))
	dsl[0] = *d
	c.data = dsl
	return ParameterConfigfileChunkType{c}
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
