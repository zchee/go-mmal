// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_queue.h>
*/
import "C"

type QueueType struct {
	c *C.MMAL_QUEUE_T
}

func QueueCreate() QueueType {
	return QueueType{C.mmal_queue_create()}
}

func QueuePut(queue *QueueType, buffer BufferHeader) {
	C.mmal_queue_put(queue.c, buffer.c)
}

func QueuePutBack(queue *QueueType, buffer BufferHeader) {
	C.mmal_queue_put_back(queue.c, buffer.c)
}

func QueueGet(queue *QueueType) BufferHeader {
	return BufferHeader{C.mmal_queue_get(queue.c)}
}

func QueueWait(queue *QueueType) BufferHeader {
	return BufferHeader{C.mmal_queue_wait(queue.c)}
}

type VcosUnsigned C.VCOS_UNSIGNED

func QueueTimedWait(queue *QueueType, timeout VcosUnsigned) BufferHeader {
	return BufferHeader{C.mmal_queue_timedwait(queue.c, C.VCOS_UNSIGNED(timeout))}
}

func QueueLength(queue *QueueType) uint {
	return uint(C.mmal_queue_length(queue.c))
}

func QueueDestroy(queue *QueueType) {
	C.mmal_queue_destroy(queue.c)
}
