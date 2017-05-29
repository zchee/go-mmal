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
//   https://github.com/Terminus-IMRC/mmal_list_supported_encodings/blob/2da6a0a40222795149863f73c14aa0b906907694/mmal_list_supported_encodings.c
// Thanks @Terminus-IMRC for publish simple source.

package main

import "C"
import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
	"unsafe"

	mmal "github.com/zchee/go-mmal"
	"github.com/zchee/go-mmal/bcmhost"
)

func init() {
	log.SetFlags(log.Lshortfile)
	log.SetPrefix("list-encoding: ")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "List supported encodings of MMAL components.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "\t%s <COMPONENT_NAME>\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example:\n\n")
		fmt.Fprintf(os.Stderr, "\t%s vc.ril.camera", os.Args[0])
	}
}

func main() {
	flag.Parse()
	componentName := flag.Arg(0)
	if componentName == "" {
		flag.Usage()
	}

	bcmhost.Init()
	mmal.LoggingInit()
	mmal.LoggingSetLevel(mmal.LogLevelTrace)

	var comp mmal.ComponentType
	if err := mmal.ComponentCreate(componentName, &comp); err != nil {
		log.Fatal(err)
	}

	fmt.Print("* Control port\n")
	encs, err := listEncodings(comp.Control())
	if err != nil && err != mmal.ErrInVal {
		log.Fatal(err)
	}
	fmt.Println(encs)

	for i := uint32(0); i < comp.InputNum(); i++ {
		fmt.Printf("\n* Input port %d\n", i)
		input := comp.Input()
		i := mmal.Port(*input[i])
		encs, err := listEncodings(&i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(encs)
	}

	for i := uint32(0); i < comp.OutputNum(); i++ {
		fmt.Printf("\n* Input port %d\n", i)
		output := comp.Output()
		encs, err := listEncodings(&output)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(encs)
	}
}

func listEncodings(port *mmal.Port) ([]string, error) {
	fmt.Printf("\t* Default encoding:      %s\n", mmal.FourccToString(uint32(port.Format().Encoding())))

	hdrp, err := mmal.PortParameterAllocGet(port, uint32(mmal.ParameterSupportedEncodings), 0)
	if err != nil {
		return nil, err
	}

	encodings := []string{}
	buf := C.GoBytes(unsafe.Pointer(&hdrp), C.int(unsafe.Sizeof(&hdrp)))
	// var enc uint32
	// binary.Read(bytes.NewReader(buf), binary.LittleEndian, enc)
	// fmt.Println(buf)
	// buf := uintptr(unsafe.Pointer(&hdrp)) + unsafe.Sizeof(mmal.ParameterHeader{})
	j := (hdrp.Size() - uint32(unsafe.Sizeof(mmal.ParameterHeader{}))) / 4
	for i := uint32(0); i < j; i++ {
		var enc uint32
		binary.Read(bytes.NewReader(buf[i*uint32(len(buf))/j:i+1*uint32(len(buf))/j]), binary.LittleEndian, enc)
		encodings = append(encodings, fmt.Sprintf("\t* Supported encoding %d: %s\n", i, mmal.FourccToString(enc)))
	}
	return encodings, nil
}
