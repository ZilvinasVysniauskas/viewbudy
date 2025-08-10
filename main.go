package main

/*
#cgo CFLAGS: -x objective-c -fmodules -fobjc-arc
#cgo LDFLAGS: -L${SRCDIR} -lplayer
#include <stdlib.h>

// Add the typedef here so Go knows about it
typedef void (*CLogCallback)(const char*);

#include "libplayer.h"
*/
import "C"

import (
	"fmt"
	"log"
	"unsafe"
)

type CLogCallback func(*C.char)

//export logMessage
func logMessage(cmsg *C.char) {
	msg := C.GoString(cmsg)
	fmt.Println("[Objective-C]", msg)
}

func main() {
	err := play("/Users/zilvis/projects/viewbudy/testvideo.mp4")
	if err != nil {
		log.Fatal(err)
	}

	select {} // Keep app running
}

func play(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	C.playVideo(cpath)
	return nil
}
