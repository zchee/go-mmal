// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_pool.h>
*/
import "C"
import (
	"unsafe"
)

type PoolType struct {
	c *C.MMAL_POOL_T
}

func (p PoolType) Queue() QueueType {
	return QueueType{p.c.queue}
}

func (p PoolType) HeadersNum() uint32 {
	return uint32(p.c.headers_num)
}

func (p PoolType) Header() BufferHeader {
	return BufferHeader{*p.c.header}
}

type PoolAllocatorAllocType C.mmal_pool_allocator_alloc_t

type PoolAllocatorFreeType C.mmal_pool_allocator_free_t

func PoolCreate(header uint, payloadSize uint32) PoolType {
	return PoolType{C.mmal_pool_create(C.uint(header), C.uint32_t(payloadSize))}
}

func PoolCreateWithAllocator(header uint, payloadSize uint32, allocatorContext unsafe.Pointer, allocatorAlloc PoolAllocatorAllocType, allocatorFree PoolAllocatorFreeType) PoolType {
	return PoolType{C.mmal_pool_create_with_allocator(C.uint(header), C.uint32_t(payloadSize), allocatorContext, C.mmal_pool_allocator_alloc_t(allocatorAlloc), C.mmal_pool_allocator_free_t(allocatorFree))}
}

func PoolDestroy(pool PoolType) {
	C.mmal_pool_destroy(pool.c)
}

func PoolResize(pool PoolType, header uint, payloadSize uint32) Status {
	return Status(C.mmal_pool_resize(pool.c, C.uint(header), C.uint32_t(payloadSize)))
}

type PoolBHCBType C.MMAL_POOL_BH_CB_T

func PoolCallbackSet(pool PoolType, cb PoolBHCBType, userdata unsafe.Pointer) {
	C.mmal_pool_callback_set(pool.c, cb, userdata)
}

func PoolPreReleaseCallbackSet(pool PoolType, cb PoolBHCBType, userdata unsafe.Pointer) {
	C.mmal_pool_pre_release_callback_set(pool.c, cb, userdata)
}
