package main

/*
#cgo CFLAGS: -x objective-c -fmodules -fobjc-arc
#cgo LDFLAGS: -L${SRCDIR} -lplayer
#include <stdlib.h>

// Add the typedef here so Go knows about it
typedef void (*CLogCallback)(const char*);

#include "../obj-c-interface/libplayer.h"
*/
import "C"

import (
	"fmt"
	"log"
	"unsafe"
)

//export logMessage
func logMessage(cmsg *C.char) {
	msg := C.GoString(cmsg)
	fmt.Println("[Objective-C]", msg)
}

func main() {
	fmt.Println("started go app")
	err := play("/Users/zilvis/projects/viewbudy/go_app/testvideo.mp4")
	if err != nil {
		log.Fatal(err)
	}

	select {} // Keep app running
}

func play(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	C.playVideo(cpath)

	// for i := 0; i < 300; i++ { // Run events for ~3 seconds
	// 	C.processEvents()
	// 	time.Sleep(10 * time.Millisecond)
	// }

	C.pauseVideo()

	return nil
}
